package v1

import (
	"jianji-server/model/request"
	"jianji-server/utils"

	"github.com/gin-gonic/gin"
)

func SetupPostRoutes(router *gin.RouterGroup) {
	TagsRouterGroup := router.Group("/posts")
	{
		TagsRouterGroup.POST("/list", ValidateCommon.AuthRequire(), PostsApi.List)
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
