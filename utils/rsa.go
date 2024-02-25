package utils

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type RSAKeyPair struct {
	privateKey *rsa.PrivateKey
	publicKey  *pem.Block
}

func GenerateRSAKeyPair(bits int) (*bytes.Buffer, *bytes.Buffer, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}

	privateKeyPem, publicKeyPEM, err2 := ConvertRSAToPEMFormat(privateKey)
	if err2 != nil {
		return nil, nil, err2
	}

	return privateKeyPem, publicKeyPEM, nil
}

func ConvertRSAToPEMFormat(privateKey *rsa.PrivateKey) (*bytes.Buffer, *bytes.Buffer, error) {
	var private, public bytes.Buffer
	// 将私钥编码为PEM格式
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	err := pem.Encode(&private, privateKeyPEM)
	if err != nil {
		return nil, nil, err
	}

	// 将公钥编码为PEM格式
	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
	}
	err2 := pem.Encode(&public, publicKeyPEM)
	if err2 != nil {
		return nil, nil, err2
	}

	return &private, &public, nil
}

func BytesToPrivateKey(privateKeyBytes []byte) (*rsa.PrivateKey, error) {
	// 解码 PEM 格式的私钥
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		return nil, errors.New("无效的 PEM 数据")
	}
	if block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("无效的 RSA 私钥")
	}

	// 解析私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func RSADecryptData(privateKey string, encryptedData string) ([]byte, error) {
	key, err := BytesToPrivateKey([]byte(privateKey))
	if err != nil {
		return nil, err
	}

	decodeString, err2 := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err2
	}

	return rsa.DecryptPKCS1v15(rand.Reader, key, decodeString)
}

// CachePrivateKeyPEM 将私钥根据jwt保存到session
func CachePrivateKeyPEM(c *gin.Context, privateKeyPEM string) error {
	session := sessions.Default(c)
	session.Set(SessionPrivateKeyPEM, privateKeyPEM)
	return session.Save()
}

// GetCachedPrivateKeyPEM 根据jwt获取私钥
func GetCachedPrivateKeyPEM(c *gin.Context) (string, bool) {
	session := sessions.Default(c)
	value := session.Get(SessionPrivateKeyPEM)
	if value == nil {
		return "", false
	} else {
		return value.(string), true
	}
}
