package v1

import (
	"jianji-server/utils/r"

	"github.com/gin-gonic/gin"
)

type Tag struct {
}

func (*Tag) List(c *gin.Context) {
	var code, message, data = TagService.List(c)
	r.OkJsonResult(c, code, message, data)
}

func (*Tag) Create(c *gin.Context) {
	var code, message, data = TagService.Create(c)
	r.OkJsonResult(c, code, message, data)
}

func (*Tag) Update(c *gin.Context) {
	var code, message, data = TagService.Update(c)
	r.OkJsonResult(c, code, message, data)
}

func (*Tag) Delete(c *gin.Context) {
	var code, message, data = TagService.Delete(c)
	r.OkJsonResult(c, code, message, data)
}
