package v1

import (
	"memo-server/model/request"
	"memo-server/utils"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.RouterGroup) {
	UserRouterGroup := router.Group("user")
	{
		UserRouterGroup.POST("/login")
		UserRouterGroup.POST(
			"/signup",
			utils.BindRequestParams[request.Signup],
			ValidateUser.Signup(),
			UserApi.Signup,
		)
		UserRouterGroup.GET("/profile")
	}

}
