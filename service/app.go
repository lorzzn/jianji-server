package service

import (
	"memo-server/utils"
	"memo-server/utils/r"

	"github.com/gin-gonic/gin"
)

type App struct{}

func (*App) GetPublicKey(c *gin.Context) (code int, message string, data any) {
	// 生成 rsa
	privateKeyPem, publicKeyPEM, err1 := utils.GenerateRSAKeyPair(2048)
	if err1 != nil {
		code = r.APP_CREATERSA_FAILED
		return
	}

	//缓存私钥
	err2 := utils.CachePrivateKeyPEM(c, privateKeyPem.Bytes)
	if err2 != nil {
		code = r.APP_SAVERSA_FAILED
		return
	}

	code = r.OK
	data = gin.H{
		"publicKey": publicKeyPEM.Bytes,
	}
	return
}
