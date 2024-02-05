package main

import (
	"log"
	"memo-server/config"
	_ "memo-server/config"
	"memo-server/middleware"
	"memo-server/routes"
	"memo-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

func main() {
	//初始化数据库
	utils.SetupDB()
	//日志
	utils.SetupLogger()

	gin.SetMode(config.Server.Mode)
	engine := gin.Default()

	engine.Use(middleware.CORS())

	routes.SetApiRoutes(engine)
	routes.SetupTestRoutes(engine)

	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "server is running!")
	})

	var port = ":8080"
	if config.Server.Port != "" {
		port = ":" + config.Server.Port
	}
	g.Go(func() error {
		return engine.Run(port)
	})

	if err := g.Wait(); err != nil {
		log.Panicln(err)
	}
}
