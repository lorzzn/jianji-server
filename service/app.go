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
	record, err := utils.ClientPublicIP(c)
	location := gin.H{}
	if err == nil {
		location = gin.H{
			"country": record.Country.Names["en"],
			"city":    record.City.Names["en"],
		}
	}
	data = gin.H{
		"time":     time.Now().UnixNano(),
		"location": location,
	}
	code = r.OK
	return
}
