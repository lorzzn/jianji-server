package service

import (
	"jianji-server/utils"
	"jianji-server/utils/r"
	"time"

	"github.com/gin-gonic/gin"
)

type App struct{}

func (*App) GetPublicKey(c *gin.Context) (code int, message string, data any) {
	// 生成 rsa
	rsaPrivate, rsaPublic, err := utils.GenerateRSAKeyPair(2048)
	if err != nil {
		code = r.APP_CREATERSA_FAILED
		return
	}

	//缓存私钥
	err2 := utils.CachePrivateKeyPEM(c, rsaPrivate.String())
	if err2 != nil {
		code = r.APP_SAVERSA_FAILED
		return
	}

	code = r.OK
	data = gin.H{
		"publicKey": rsaPublic.String(),
	}
	return
}

func (*App) GetAppConfig(c *gin.Context) (code int, message string, data any) {
	ip := utils.GetClientIP(c)
	location := gin.H{
		"country": "",
		"city":    "",
	}

	record, err := utils.GetIPGeoRecord(ip)
	if err == nil {
		location = gin.H{
			"country": record.Country.Names["en"],
			"city":    record.City.Names["en"],
		}
	}
	sessionId, _ := c.Get("SessionId")
	data = gin.H{
		"time":       time.Now().UnixMilli(),
		"location":   location,
		"session_id": sessionId,
	}
	code = r.OK
	return
}
