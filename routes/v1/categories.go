package v1

import (
	"jianji-server/model/request"
	"jianji-server/utils"

	"github.com/gin-gonic/gin"
)

func SetupCategoriesRoutes(router *gin.RouterGroup) {
	CategoriesRouterGroup := router.Group("/categories")
	{
		CategoriesRouterGroup.POST("/load", CategoriesApi.Load)
		CategoriesRouterGroup.POST(
			"/create",
			utils.BindRequestParams[request.CreateCategories],
			ValidateCategories.CreateCategories(),
			CategoriesApi.Create,
		)
		CategoriesRouterGroup.POST(
			"/delete",
			utils.BindRequestParams[request.DeleteCategories],
			ValidateCategories.DeleteCategories(),
			CategoriesApi.Delete,
		)
	}
}
