package common

import (
	"jianji-server/utils/r"

	"github.com/gin-gonic/gin"
)

type App struct {
}

func (*App) GetPublicKey(c *gin.Context) {
	var code, message, data = AppService.GetPublicKey(c)
	r.OkJsonResult(c, code, message, data)
}

func (*App) GetAppConfig(c *gin.Context) {
	var code, message, data = AppService.GetAppConfig(c)
	r.OkJsonResult(c, code, message, data)
}
