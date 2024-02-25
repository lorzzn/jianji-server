package utils

import (
	"jianji-server/config"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var RDB redis.Client
var RedisGlobalContext = context.Background()

func SetupRedis() {
	// 创建 Redis 客户端连接
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,     // Redis 服务器地址
		Password: config.Redis.Password, // Redis 访问密码，如果没有设置密码则为空字符串
		DB:       config.Redis.DB,       // 选择使用的数据库，默认为0
	})

	RDB = *rdb
}
