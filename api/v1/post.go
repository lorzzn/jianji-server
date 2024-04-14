package v1

import (
	"jianji-server/utils/r"

	"github.com/gin-gonic/gin"
)

type Posts struct {
}

func (*Posts) List(c *gin.Context) {
	var code, message, data = PostsService.List(c)
	r.OkJsonResult(c, code, message, data)
}

func (*Posts) Get(c *gin.Context) {
	var code, message, data = PostsService.Get(c)
	r.OkJsonResult(c, code, message, data)
}

func (*Posts) Create(c *gin.Context) {
	var code, message, data = PostsService.Create(c)
	r.OkJsonResult(c, code, message, data)
}

func (*Posts) Update(c *gin.Context) {
	var code, message, data = PostsService.Update(c)
	r.OkJsonResult(c, code, message, data)
}

func (*Posts) Delete(c *gin.Context) {
	var code, message, data = PostsService.Delete(c)
	r.OkJsonResult(c, code, message, data)
}
