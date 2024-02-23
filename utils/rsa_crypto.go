package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"github.com/gin-gonic/gin"
)

type RSAKeyPair struct {
	privateKey *rsa.PrivateKey
	publicKey  *pem.Block
}

func GenerateRSAKeyPair(bits int) (*pem.Block, *pem.Block, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}

	privateKeyPem, publicKeyPEM := ConvertRSAToPEMFormat(privateKey)

	return privateKeyPem, publicKeyPEM, nil
}

func ConvertRSAToPEMFormat(privateKey *rsa.PrivateKey) (*pem.Block, *pem.Block) {
	// 将私钥编码为PEM格式
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	// 将公钥编码为PEM格式
	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
	}
	return privateKeyPEM, publicKeyPEM
}

func createRedisKey(jwtId string) string {
	return "rsa:" + jwtId
}

func getGinContextJwtId(c *gin.Context) (string, error) {

	jwtId, existed := c.Get("JwtId")
	if !existed {
		return "", errors.New("jwtId 获取错误")
	}
	jwtIdStr, ok := jwtId.(string)
	if !ok {
		return "", errors.New("jwtId 格式错误")
	}
	return jwtIdStr, nil
}

// CachePrivateKeyPEM 将私钥根据jwt保存到redis
func CachePrivateKeyPEM(c *gin.Context, privateKeyPEM []byte) error {
	jwtId, err := getGinContextJwtId(c)
	if err != nil {
		return err
	}

	//redis 储存
	return RDB.Set(RedisGlobalContext, createRedisKey(jwtId), string(privateKeyPEM), AccessTokenExpireDuration).Err()
}

// GetCachedPrivateKeyPEM 根据jwt获取私钥
func GetCachedPrivateKeyPEM(c *gin.Context) (string, error) {
	jwtId, err1 := getGinContextJwtId(c)
	if err1 != nil {
		return "", err1
	}

	//redis 读取
	value, err2 := RDB.Get(RedisGlobalContext, createRedisKey(jwtId)).Result()
	return value, err2
}
