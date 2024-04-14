package v1

import (
	"jianji-server/model/request"
	"jianji-server/utils"

	"github.com/gin-gonic/gin"
)

func SetupPostRoutes(router *gin.RouterGroup) {
	TagsRouterGroup := router.Group("/posts")
	{
		TagsRouterGroup.POST(
			"/list",
			ValidateCommon.AuthRequire(),
			utils.BindRequestParams[request.ListPost],
			PostsApi.List,
		)
		TagsRouterGroup.POST(
			"/get",
			ValidateCommon.AuthRequire(),
			utils.BindRequestParams[request.GetPost],
			PostsApi.Get,
		)
		TagsRouterGroup.POST(
			"/create",
			ValidateCommon.AuthRequire(),
			utils.BindRequestParams[request.CreatePost],
			PostsApi.Create,
		)
		TagsRouterGroup.POST(
			"/update",
			ValidateCommon.AuthRequire(),
			utils.BindRequestParams[request.UpdatePost],
			PostsApi.Update,
		)
		TagsRouterGroup.POST(
			"/delete",
			ValidateCommon.AuthRequire(),
			utils.BindRequestParams[request.DeletePost],
			PostsApi.Delete,
		)
	}
}
