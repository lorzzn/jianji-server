package main

import (
	"jianji-server/config"
	_ "jianji-server/config"
	"jianji-server/middleware"
	"jianji-server/routes"
	"jianji-server/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

func main() {
	//日志
	utils.SetupLogger()
	//cron
	utils.SetupCron()
	if utils.ConsoleConfirm("是否更新geoip数据库？", 2) {
		//更新geoip 数据库
		utils.SetupMmdb()
	}
	//初始化数据库
	utils.SetupDB()
	//redis
	utils.SetupRedis()

	//sessionManager
	utils.SetupSessionManager()

	defer func(RDB *redis.Client) {
		err := RDB.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(&utils.RDB)

	//应用配置
	gin.SetMode(config.Server.Mode)
	engine := gin.Default()

	//应用中间件
	engine.Use(middleware.CORS())
	engine.Use(middleware.SessionMiddleware())
	engine.Use(middleware.SessionIdMiddleware())
	engine.Use(middleware.JWTAuthMiddleware())
	engine.Use(middleware.RequestIdMiddleWare())
	engine.Use(middleware.DecryptMiddleware())
	engine.Use(middleware.TraceIdMiddleware())
	engine.Use(middleware.TxQueryMiddleware())

	// log 相关中间件
	engine.Use(middleware.ResponseLoggerMiddleware())
	engine.Use(middleware.ErrorLoggerMiddleware())
	engine.Use(middleware.GinzapLoggerMiddleware())

	//应用路由
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
