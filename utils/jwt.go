package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// CustomClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个UserId字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type CustomClaims struct {
	UserId   uint64    `json:"user_id"`
	UserUUID uuid.UUID `json:"user_uuid"`
	JwtUUID  uuid.UUID `json:"jwt_id"`
	jwt.StandardClaims
}

var (
	// 定义Secret 用于加密的字符串
	jwtSecret = []byte("note car ball blue cat line")

	// 拉黑的 jwt 在 redis 里的集合
	blacklistKey = "jwt_blacklist"

	// AccessTokenExpireDuration 定义Token的过期时间
	AccessTokenExpireDuration = time.Hour * 24 // access_token 过期时间
	// RefreshTokenExpireDuration 定义RefreshToken的过期时间
	RefreshTokenExpireDuration = time.Hour * 24 * 7 // refresh_token 过期时间
)

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return jwtSecret, nil
}

// GenToken 生成JWT 生成 access_token 和 refresh_token
func GenToken(userid uint64, userUUID uuid.UUID) (accessToken, refreshToken string, err error) {
	// 创建一个我们自己的声明
	c := CustomClaims{
		userid, // 自定义字段
		userUUID,
		uuid.New(),
		jwt.StandardClaims{ // JWT规定的7个官方字段
			ExpiresAt: time.Now().Add(AccessTokenExpireDuration).Unix(), // 过期时间
			Issuer:    "server",                                         // 签发人
		},
	}
	// 加密并获得完整的编码后的字符串token
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(jwtSecret)

	// refresh token 不需要存任何自定义数据
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(RefreshTokenExpireDuration).Unix(), // 过期时间
		Issuer:    "server",                                          // 签发人
	}).SignedString(jwtSecret)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (claims *CustomClaims, err error) {
	// 解析token
	var token *jwt.Token
	claims = new(CustomClaims)
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
func RefreshToken(accessToken, refreshToken string) (newAccessToken, newRefreshToken string, err error) {
	newAccessToken = accessToken
	newRefreshToken = refreshToken

	// refresh token无效直接返回
	if _, err = jwt.Parse(refreshToken, keyFunc); err != nil {
		return
	}

	// 从旧access token中解析出claims数据	解析出payload负载信息
	var claims CustomClaims
	_, tokenErr := jwt.ParseWithClaims(accessToken, &claims, keyFunc)

	//检查token是否已经拉黑
	blacklisted, err := CheckJWTIsBlacklisted(claims.JwtUUID.String())
	if err != nil {
		return
	}

	if blacklisted {
		err = errors.New("登录已失效")
		return
	}

	// 当access token是过期错误 并且 refresh token没有过期时就创建一个新的access token
	var v *jwt.ValidationError
	var ok = errors.As(tokenErr, &v)
	if ok && v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claims.UserId, claims.UserUUID)
	}

	return
}

// AddJWTToBlacklist 模拟将令牌加入黑名单
func AddJWTToBlacklist(jwtUUID string) error {
	return RDB.Set(RedisGlobalContext, fmt.Sprintf("%s:%s", blacklistKey, jwtUUID), 1, AccessTokenExpireDuration).Err()
}

// CheckJWTIsBlacklisted 检查令牌是否在黑名单中
func CheckJWTIsBlacklisted(jwtUUID string) (bool, error) {
	result, err := RDB.Exists(RedisGlobalContext, fmt.Sprintf("%s:%s", blacklistKey, jwtUUID)).Result()
	return result == 1, err
}
