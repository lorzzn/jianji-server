package routes

import (
	"memo-server/routes/common"
	v1 "memo-server/routes/v1"

	"github.com/gin-gonic/gin"
)

func SetApiRoutes(engine *gin.Engine) {
	CommonRouteGroup := engine.Group("/api/common")
	{
		common.SetCommonAppRoutes(CommonRouteGroup)
	}

	ApiV1RouteGroup := engine.Group("/api/v1")
	{
		v1.SetupUserRoutes(ApiV1RouteGroup)
	}
}
