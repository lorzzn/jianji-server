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

// Package main provides a simple HTTP API for memo.
//
//     Schemes: http
//     BasePath: /api
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta

func main() {
	//初始化数据库
	utils.SetupDB()
	//日志
	utils.SetupLogger()

	gin.SetMode(config.Server.Mode)
	router := gin.Default()

	router.Use(middleware.CORS())

	routes.SetApiRoutes(router)

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "server is running!")
	})

	var port = ":8080"
	if config.Server.Port != "" {
		port = ":" + config.Server.Port
	}
	g.Go(func() error {
		return router.Run(port)
	})

	if err := g.Wait(); err != nil {
		log.Panicln(err)
	}
}
