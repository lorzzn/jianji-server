package r

import (
	"encoding/json"
	"net/http"

	"github.com/Rican7/conjson"
	"github.com/Rican7/conjson/transform"
	"github.com/gin-gonic/gin"
)

// Response 响应结构体
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func JsonResult(c *gin.Context, httpCode, code int, msg string, data any) {
	marshaller := conjson.NewMarshaler(Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}, transform.CamelCaseKeys(false))

	data2, _ := json.Marshal(marshaller)

	c.Header("Content-Type", "application/json")
	c.String(httpCode, string(data2))
}

func OkJsonResult(c *gin.Context, code int, message string, data any) {
	if message == "" {
		message = GetCodeMsg(code)
	}
	JsonResult(c, http.StatusOK, code, message, data)
}
