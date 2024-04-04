package v1

import (
	"jianji-server/model/request"
	"jianji-server/utils"

	"github.com/gin-gonic/gin"
)

func SetupTagRoutes(router *gin.RouterGroup) {
	TagRouterGroup := router.Group("/tag")
	{
		TagRouterGroup.POST("/list", ValidateCommon.AuthRequire(), TagApi.List)
		TagRouterGroup.POST(
			"/create",
			ValidateCommon.AuthRequire(),
			utils.BindRequestParams[request.CreateTag],
			ValidateTag.CreateTag(),
			TagApi.Create,
		)
		TagRouterGroup.POST(
			"/update",
			ValidateCommon.AuthRequire(),
			utils.BindRequestParams[request.UpdateTag],
			TagApi.Update,
		)
		TagRouterGroup.POST(
			"/delete",
			ValidateCommon.AuthRequire(),
			utils.BindRequestParams[request.DeleteTagBatch],
			ValidateTag.DeleteTag(),
			TagApi.Delete,
		)
	}
}
