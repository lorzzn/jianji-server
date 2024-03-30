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
	//和公共Logger写入同一个log文件
	logger := zap.New(zapcore.NewCore(utils.LoggerEncoder, utils.LoggerFileSyncer, utils.LoggerLevelEnabler))
	return ginzap.GinzapWithConfig(logger, &ginzap.Config{
		UTC:        true,
		TimeFormat: time.RFC3339,
	})
}
