package common

import "github.com/gin-gonic/gin"

func SetCommonAppRoutes(router *gin.RouterGroup) {
	CommonAppRouterGroup := router.Group("app")
	{
		CommonAppRouterGroup.GET("/public-key", AppApi.GetPublicKey)
		CommonAppRouterGroup.GET("/info", AppApi.GetAppInfo)
	}
}
