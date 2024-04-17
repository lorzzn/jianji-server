package v1

import (
	"jianji-server/utils/r"

	"github.com/gin-gonic/gin"
)

type Categories struct {
}

func (*Categories) List(c *gin.Context) {
	var code, message, data = CategoriesService.List(c)
	r.OkJsonResult(c, code, message, data)
}

func (*Categories) Create(c *gin.Context) {
	var code, message, data = CategoriesService.Create(c)
	r.OkJsonResult(c, code, message, data)
}

func (*Categories) Update(c *gin.Context) {
	var code, message, data = CategoriesService.Update(c)
	r.OkJsonResult(c, code, message, data)
}

func (*Categories) Delete(c *gin.Context) {
	var code, message, data = CategoriesService.Delete(c)
	r.OkJsonResult(c, code, message, data)
}

func (*Categories) CategoryStatistics(c *gin.Context) {
	var code, message, data = CategoriesService.CategoryStatistics(c)
	r.OkJsonResult(c, code, message, data)
}
