package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个UserId字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserId any `json:"user_id"`
	JwtId  any `json:"jwt_id"`
	jwt.StandardClaims
}

// 定义Secret 用于加密的字符串
var mySecret = []byte("note car ball blue cat line")

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return mySecret, nil
}

// TokenExpireDuration 定义JWT的过期时间
const TokenExpireDuration = time.Hour * 24

// AccessTokenExpireDuration 定义JWT的过期时间
const AccessTokenExpireDuration = time.Hour * 24      // access_token 过期时间
const RefreshTokenExpireDuration = time.Hour * 24 * 7 // refresh_token 过期时间

// GenToken 生成JWT 生成 access_token 和 refresh_token
func GenToken(userid any) (accessToken, refreshToken string, err error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		userid, // 自定义字段
		uuid.New(),
		jwt.StandardClaims{ // JWT规定的7个官方字段
			ExpiresAt: time.Now().Add(AccessTokenExpireDuration).Unix(), // 过期时间
			Issuer:    "server",                                         // 签发人
		},
	}
	// 加密并获得完整的编码后的字符串token
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret)

	// refresh token 不需要存任何自定义数据
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(RefreshTokenExpireDuration).Unix(), // 过期时间
		Issuer:    "server",                                          // 签发人
	}).SignedString(mySecret)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (claims *MyClaims, err error) {
	// 解析token
	var token *jwt.Token
	claims = new(MyClaims)
	token, err = jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		return
	}
	if !token.Valid { // 校验token
		err = errors.New("invalid token")
	}
	return
}

// RefreshToken 刷新AccessToken
func RefreshToken(accessToken, refreshToken string) (newAToken, newRToken string, err error) {
	// refresh token无效直接返回
	if _, err = jwt.Parse(refreshToken, keyFunc); err != nil {
		return
	}

	// 从旧access token中解析出claims数据	解析出payload负载信息
	var claims MyClaims
	_, err = jwt.ParseWithClaims(accessToken, &claims, keyFunc)
	v, _ := err.(*jwt.ValidationError)

	// 当access token是过期错误 并且 refresh token没有过期时就创建一个新的access token
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claims.UserId)
	}
	return
}
