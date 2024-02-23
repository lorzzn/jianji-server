package v1

import (
	"memo-server/model/request"
	"memo-server/utils"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.RouterGroup) {
	UserRouterGroup := router.Group("user")
	{
		UserRouterGroup.POST(
			"/login",
			utils.BindRequestParams[request.Login],
			ValidateUser.Login(),
			UserApi.Login,
		)
		UserRouterGroup.GET("/profile")
	}

}
