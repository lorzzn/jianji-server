package v1

import (
	"jianji-server/model/request"
	"jianji-server/utils"

	"github.com/gin-gonic/gin"
)

func SetupTagRoutes(router *gin.RouterGroup) {
	TagsRouterGroup := router.Group("/tags")
	{
		TagsRouterGroup.POST("/list", ValidateCommon.AuthRequire(), TagsApi.List)
		TagsRouterGroup.POST(
			"/create",
			ValidateCommon.AuthRequire(),
			utils.BindRequestParams[request.CreateTags],
			ValidateTags.CreateTags(),
			TagsApi.Create,
		)
		TagsRouterGroup.POST(
			"/update",
			ValidateCommon.AuthRequire(),
			utils.BindRequestParams[request.UpdateTags],
			TagsApi.Update,
		)
		TagsRouterGroup.POST(
			"/delete",
			ValidateCommon.AuthRequire(),
			utils.BindRequestParams[request.DeleteTags],
			ValidateTags.DeleteTags(),
			TagsApi.Delete,
		)
		TagsRouterGroup.POST(
			"/statistics",
			ValidateCommon.AuthRequire(),
			utils.BindRequestParams[request.TagStatistics],
			ValidateTags.TagStatistics(),
			TagsApi.TagStatistics,
		)
	}
}
