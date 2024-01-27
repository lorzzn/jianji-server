package r

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 响应结构体
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Result(c *gin.Context, httpCode, code int, msg string, data any) {
	c.JSON(httpCode, Response{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}

func OkResult(c *gin.Context, code int, message string, data any) {
	if message == "" {
		message = GetCodeMsg(code)
	}
	Result(c, http.StatusOK, code, message, data)
}
