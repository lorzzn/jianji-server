package routes

import (
	"jianji-server/utils/r"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func SetupTestRoutes(engine *gin.Engine) {
	TestRouteGroup := engine.Group("/test")
	{
		TestRouteGroup.POST("/request", func(c *gin.Context) {
			dump, _ := httputil.DumpRequest(c.Request, true)
			c.JSON(200, r.Response{
				Code:    0,
				Message: "success",
				Data:    string(dump),
			})
		})
	}
}
