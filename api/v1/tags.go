package v1

import (
	"jianji-server/utils/r"

	"github.com/gin-gonic/gin"
)

type Tags struct {
}

func (*Tags) List(c *gin.Context) {
	var code, message, data = TagsService.List(c)
	r.OkJsonResult(c, code, message, data)
}

func (*Tags) Create(c *gin.Context) {
	var code, message, data = TagsService.Create(c)
	r.OkJsonResult(c, code, message, data)
}

func (*Tags) Update(c *gin.Context) {
	var code, message, data = TagsService.Update(c)
	r.OkJsonResult(c, code, message, data)
}

func (*Tags) Delete(c *gin.Context) {
	var code, message, data = TagsService.Delete(c)
	r.OkJsonResult(c, code, message, data)
}

func (*Tags) TagStatistics(c *gin.Context) {
	var code, message, data = TagsService.TagStatistics(c)
	r.OkJsonResult(c, code, message, data)
}
