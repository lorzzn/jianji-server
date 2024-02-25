package utils

import (
	"encoding/json"
	"errors"
	"memo-server/utils/r"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const ContextRequestParams = "RequestParams"
const ContextDecryptedParams = "DecryptedParams"

func GetRequestParams[T any](c *gin.Context) (params T, ok bool) {
	var value any
	value, ok = c.Get(ContextRequestParams)
	params = value.(T)
	return
}

func BindRequestParams[T any](c *gin.Context) {
	var params T
	var err error
	//判断是否是加密请求
	isEncrypted := c.Request.Header.Get("Encrypted")
	if isEncrypted == "true" {
		decrypted, ok := c.Get(ContextDecryptedParams)
		if ok {
			err = json.Unmarshal(decrypted.([]byte), &params)
		} else {
			err = errors.New(ContextDecryptedParams + " 获取失败")
		}
	} else {
		err = c.ShouldBind(&params)
	}
	if err != nil {
		Logger.Error("BindRequestParams", zap.Error(err))
		r.OkJsonResult(c, r.ERROR_BAD_PARAM, "请求参数获取失败", nil)
		c.Abort()
		return
	}
	c.Set(ContextRequestParams, params)
	c.Next()
}
