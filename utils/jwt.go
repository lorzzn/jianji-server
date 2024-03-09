package utils

import (
	"errors"
	"fmt"
	"jianji-server/entity"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CustomClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个UserId字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type CustomClaims struct {
	UserId            uint64    `json:"user_id"`
	UserUUID          uuid.UUID `json:"user_uuid"`
	TokenUUID         uuid.UUID `json:"token_uuid"`
	ClientFingerprint string    `json:"client_fingerprint"`
	jwt.StandardClaims
}

type TokenData struct {
	Token        string
	RefreshToken string
	TokenUUID    uuid.UUID
	ExpiresAt    time.Time
}

var (
	// 定义Secret 用于加密的字符串
	jwtSecret = []byte("note car ball blue cat line")

	// 拉黑的 jwt 在 redis 里的集合
	blacklistKey = "jwt_blacklist"

	// TokenExpireDuration 定义Token的过期时间
	TokenExpireDuration = time.Hour * 24 // access_token 过期时间
	// RefreshTokenExpireDuration 定义RefreshToken的过期时间
	RefreshTokenExpireDuration = time.Hour * 24 * 7 // refresh_token 过期时间
)

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return jwtSecret, nil
}

// GenToken 生成JWT 生成 access_token 和 refresh_token
func GenToken(c *gin.Context, userID uint64, userUUID uuid.UUID, fingerprint string) (*TokenData, error) {
	// 创建一个我们自己的声明
	tokenUUID := uuid.New()
	expiresAt := time.Now().Add(TokenExpireDuration)
	claims := CustomClaims{
		userID, // 自定义字段
		userUUID,
		tokenUUID,
		fingerprint,
		jwt.StandardClaims{ // JWT规定的7个官方字段
			ExpiresAt: expiresAt.Unix(), // 过期时间
			Issuer:    "server",         // 签发人
		},
	}
	// 加密并获得完整的编码后的字符串token
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	// refresh token 不需要存任何自定义数据
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(RefreshTokenExpireDuration).Unix(), // 过期时间
		Issuer:    "server",                                          // 签发人
	}).SignedString(jwtSecret) // 使用指定的secret签名并获得完整的编码后的字符串token

	//将jwt token授权的详细信息储存到数据库
	userToken := &entity.UserToken{
		UserUUID:          userUUID,
		Token:             token,
		TokenUUID:         tokenUUID,
		ClientFingerprint: fingerprint,
		UserAgent:         c.Request.UserAgent(),
		ExpiresAt:         expiresAt,
	}
	ip := GetClientIP(c)
	geoRecord, err := GetIPGeoRecord(ip)
	if err == nil {
		userToken.Country = geoRecord.Country.Names["en"]
		userToken.City = geoRecord.City.Names["en"]
	}

	DB.Create(userToken)

	return &TokenData{
		Token:        token,
		RefreshToken: refreshToken,
		TokenUUID:    tokenUUID,
		ExpiresAt:    expiresAt,
	}, nil
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

// RefreshToken 刷新Token
func RefreshToken(c *gin.Context, token, refreshToken string) (*TokenData, error) {
	newToken := token
	newRefreshToken := refreshToken

	// refresh token无效，返回报错err
	if _, err := jwt.Parse(refreshToken, keyFunc); err != nil {
		return nil, err
	}

	// 从旧access token中解析出claims数据	解析出payload负载信息
	var claims CustomClaims
	_, tokenErr := jwt.ParseWithClaims(token, &claims, keyFunc)

	//检查token是否已经拉黑，存在报错返回err
	blacklisted, err := IsTokenBlacklisted(claims.TokenUUID.String())
	if err != nil {
		return nil, err
	}

	//token已经拉黑，返回报错
	if blacklisted {
		err = errors.New("登录已失效")
		return nil, err
	}

	// 当access token是过期错误 并且 refresh token没有过期时就创建一个新的access token
	var v *jwt.ValidationError
	var ok = errors.As(tokenErr, &v)
	if ok && v.Errors == jwt.ValidationErrorExpired {
		return GenToken(c, claims.UserId, claims.UserUUID, claims.ClientFingerprint)
	}

	return &TokenData{
		Token:        newToken,
		RefreshToken: newRefreshToken,
		TokenUUID:    claims.TokenUUID,
		ExpiresAt:    time.Unix(claims.ExpiresAt, 0),
	}, nil
}

// AddTokenToBlacklist 模拟将令牌加入黑名单
func AddTokenToBlacklist(tokenUUID string) {
	tUUID, err := uuid.Parse(tokenUUID)
	if err != nil {
		Logger.Error(err.Error())
	}
	err = DB.Where(&entity.UserToken{TokenUUID: tUUID}).Limit(1).Updates(&entity.UserToken{Blacklisted: true}).Error
	if err != nil {
		Logger.Error(err.Error())
	}
	err = RDB.Set(RedisGlobalContext, fmt.Sprintf("%s:%s", blacklistKey, tokenUUID), 1, TokenExpireDuration).Err()
	if err != nil {
		Logger.Error(err.Error())
	}
}

// IsTokenBlacklisted 检查令牌是否在黑名单中
func IsTokenBlacklisted(tokenUUID string) (bool, error) {
	result, err := RDB.Exists(RedisGlobalContext, fmt.Sprintf("%s:%s", blacklistKey, tokenUUID)).Result()
	return result == 1, err
}

func CleanExpiredDatabaseUserToken() {
	DB.Where("expires_at < ?", time.Now()).Updates(&entity.UserToken{Status: -1})
}
