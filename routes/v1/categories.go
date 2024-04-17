package v1

import (
	"jianji-server/model/request"
	"jianji-server/utils"

	"github.com/gin-gonic/gin"
)

func SetupCategoriesRoutes(router *gin.RouterGroup) {
	CategoriesRouterGroup := router.Group("/categories")
	{
		CategoriesRouterGroup.POST("/list", ValidateCommon.AuthRequire(), CategoriesApi.List)
		CategoriesRouterGroup.POST(
			"/create",
			ValidateCommon.AuthRequire(),
			utils.BindRequestParams[request.CreateCategories],
			ValidateCategories.CreateCategories(),
			CategoriesApi.Create,
		)
		CategoriesRouterGroup.POST(
			"/update",
			ValidateCommon.AuthRequire(),
			utils.BindRequestParams[request.UpdateCategories],
			CategoriesApi.Update,
		)
		CategoriesRouterGroup.POST(
			"/delete",
			ValidateCommon.AuthRequire(),
			utils.BindRequestParams[request.DeleteCategories],
			ValidateCategories.DeleteCategories(),
			CategoriesApi.Delete,
		)
		CategoriesRouterGroup.POST(
			"/statistics",
			ValidateCommon.AuthRequire(),
			utils.BindRequestParams[request.CategoryStatistics],
			ValidateCategories.CategoryStatistics(),
			CategoriesApi.CategoryStatistics,
		)
	}
}
