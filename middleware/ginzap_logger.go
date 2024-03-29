package middleware

import (
	"jianji-server/utils"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GinzapLoggerMiddleware() gin.HandlerFunc {
	//共用公共Logger，写入公共的log文件
	logger := utils.Logger.WithOptions(zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return zapcore.NewCore(utils.LoggerEncoder, utils.LoggerFileSyncer, utils.LoggerLevelEnabler)
	}))

	return ginzap.GinzapWithConfig(logger, &ginzap.Config{
		UTC:        true,
		TimeFormat: time.RFC3339,
	})
}
