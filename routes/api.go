package routes

import (
	v1 "memo-server/routes/v1"

	"github.com/gin-gonic/gin"
)

func SetApiRoutes(engine *gin.Engine) {
	ApiV1RouteGroup := engine.Group("/api/v1")
	{
		v1.SetupUserRoutes(ApiV1RouteGroup)
	}
}
