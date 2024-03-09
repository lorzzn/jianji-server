package utils

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"jianji-server/config"
	"net"
	"net/mail"
	"net/smtp"
	"time"
)

type EmailActiveData struct {
	State       string
	Email       string
	Password    string
	Fingerprint string
}

// SendEmail https://gist.github.com/chrisgillis/10888032
func SendEmail(receiver string, subject string, body string) error {

	from := mail.Address{Name: "简记", Address: config.Email.From}
	to := mail.Address{Address: receiver}

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := fmt.Sprintf("%s:%d", config.Email.SMTPServer, config.Email.SMTPPort)

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", config.Email.From, config.Email.Password, host)

	// TLS config
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require a ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", servername, tlsConfig)
	if err != nil {
		return err
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		return err
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		return err
	}

	if err = c.Rcpt(to.Address); err != nil {
		return err
	}

	// Data
	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return c.Quit()
}

func GetWaitingActivationEmailKey(email string) string {
	return "waiting_activation_email:" + email
}

// SendActiveEmail 发送激活邮件
func SendActiveEmail(email string, password string, fingerprint string) error {
	// 生成邮箱激活码, 将激活码根据邮箱存入redis
	state := GenerateRandomString(16)
	key := GetWaitingActivationEmailKey(email)

	b, err := json.Marshal(EmailActiveData{
		State:       state,
		Email:       email,
		Password:    password,
		Fingerprint: fingerprint,
	})
	if err != nil {
		return err
	}

	err = RDB.Set(RedisGlobalContext, key, string(b), time.Minute*30).Err()
	if err != nil {
		return err
	}

	// 发送激活链接到用户邮箱
	err = SendEmail(
		email,
		"欢迎使用 简记 — 确认注册",
		fmt.Sprintf(
			"请点击下面的链接激活你的账户（链接30分钟内有效）\n%s/active?email=%s&state=%s",
			config.Server.WebDomain,
			email,
			state,
		),
	)
	if err != nil {
		return err
	}

	return nil
}

// GetActiveEmailStateInfo 根据激活链接拿到用户注册信息
func GetActiveEmailStateInfo(email string, state string) (string, string, string, error) {
	key := GetWaitingActivationEmailKey(email)

	// 无效激活链接情况
	value, err := RDB.Get(RedisGlobalContext, key).Result()
	if err != nil {
		return "", "", "", errors.New("无效的激活链接")
	}

	// 获取后删掉redis中的state数据
	err = RDB.Del(RedisGlobalContext, key).Err()
	if err != nil {
		return "", "", "", errors.New("系统发生错误")
	}

	var data EmailActiveData
	if err = json.Unmarshal([]byte(value), &data); err != nil {
		return "", "", "", errors.New("系统发生错误")
	}

	// 用户重新获取激活链接的情况
	if data.State != state {
		return "", "", "", errors.New("激活链接已失效")
	}

	return data.Email, data.Password, data.Fingerprint, nil
}
