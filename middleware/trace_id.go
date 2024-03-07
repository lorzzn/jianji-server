package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TraceIdMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		traceID := uuid.New().String()
		c.Set("TraceId", traceID)
		c.Header("X-Trace-ID", traceID)
	}
}
