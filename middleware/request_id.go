package middleware

import (
	"memo-server/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func RequestIdMiddleWare() func(c *gin.Context) {
	return func(c *gin.Context) {
		requestId := c.Request.Header.Get("X-Request-ID")
		if requestId == "" {
			requestId = utils.GenerateRequestId()
		}
		//统一转换为小写
		requestId = strings.ToLower(requestId)

		c.Set("RequestId", requestId)
		c.Next()
		return
	}
}
