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
			UserApi.Signup,
		)
		UserRouterGroup.POST(
			"/signup",
			utils.BindRequestParams[request.Signup],
			ValidateUser.Signup(),
			UserApi.Signup,
		)
		UserRouterGroup.POST(
			"/logout",
			UserApi.Logout,
		)
		UserRouterGroup.POST(
			"/active",
			utils.BindRequestParams[request.Active],
			ValidateUser.Active(),
			UserApi.Active,
			UserApi.Login,
		)
		UserRouterGroup.POST(
			"/refresh-token",
			utils.BindRequestParams[request.RefreshToken],
			ValidateUser.RefreshToken(),
			UserApi.RefreshToken,
		)
		UserRouterGroup.POST(
			"/profile",
			UserApi.GetProfile,
		)
		UserRouterGroup.POST(
			"/edit-profile",
			utils.BindRequestParams[request.EditProfile],
			ValidateUser.EditProfile(),
			UserApi.EditProfile,
		)
	}

}
