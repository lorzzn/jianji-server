package middleware

import (
	"jianji-server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ErrorLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			//检查运行时报错
			if r := recover(); r != nil {
				err, _ := r.(error)
				utils.Logger.Panic("ErrorLogger: runtime error", zap.Error(err), zap.Stack("stack"))
			}
			// 检查gin报错
			err := c.Errors.JSON()
			if err != nil {
				utils.Logger.Panic("ErrorLogger: gin error", zap.Any("error", err), zap.Stack("stack"))
			}
		}()
		c.Next()
	}
}
