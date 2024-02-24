package common

import (
	"memo-server/utils/r"

	"github.com/gin-gonic/gin"
)

type App struct {
}

func (*App) GetPublicKey(c *gin.Context) {
	var code, message, data = AppService.GetPublicKey(c)
	r.OkJsonResult(c, code, message, data)
}

func (*App) GetAppInfo(c *gin.Context) {
	var code, message, data = AppService.GetAppInfo(c)
	r.OkJsonResult(c, code, message, data)
}
