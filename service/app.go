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

func (*App) GetAppConfig() (code int, message string, data any) {
	data = gin.H{
		"time": time.Now().UnixNano(),
	}
	code = r.OK
	return
}
