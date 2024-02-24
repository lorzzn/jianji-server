package utils

import (
	"encoding/json"
	"fmt"
	"memo-server/utils/r"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const ContextRequestParams = "RequestParams"

func GetRequestParams[T any](c *gin.Context) (params T, ok bool) {
	var value any
	value, ok = c.Get(ContextRequestParams)
	params = value.(T)
	return
}

func BindRequestParams[T any](c *gin.Context) {
	var params T
	var err error
	decryptedData, ok := c.Get("DecryptedData")
	if ok {
		err = json.Unmarshal(decryptedData.([]byte), &params)
	} else {
		err = c.ShouldBind(&params)
	}
	if err != nil {
		Logger.Error("BindRequestParams", zap.Error(err))
		r.JsonResult(c, http.StatusOK, r.ERROR_BAD_PARAM, fmt.Sprint(err), nil)
		c.Abort()
		return
	}
	c.Set(ContextRequestParams, params)
	c.Next()
}
