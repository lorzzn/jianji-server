package v1

import (
	"jianji-server/model/request"
	"jianji-server/utils"

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
