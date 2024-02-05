package routes

import (
	"memo-server/utils/r"

	"github.com/gin-gonic/gin"
)

func ReqInfo(c *gin.Context) {
	c.JSON(200, r.Response{
		Code:    0,
		Message: "req_info",
		Data:    nil,
	})
}

func SetupTestRoutes(engine *gin.Engine) {
	TestRouteGroup := engine.Group("/test")
	{
		TestRouteGroup.POST("/req_info", ReqInfo)
	}
}
