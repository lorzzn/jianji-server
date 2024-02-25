package utils

import (
	"log"
	"memo-server/config"

	"github.com/gin-gonic/contrib/sessions"
)

var SessionStore sessions.RedisStore

// 声明session key
const (
	SessionPrivateKeyPEM = "privateKeyPEM"
)

func SetupSessionManager() {
	store, err := sessions.NewRedisStore(
		10,
		config.Redis.Network,
		config.Redis.Addr,
		config.Redis.Password,
		[]byte("store"),
	)

	if err != nil {
		log.Panicln("session manager 初始化失败", err)
	}

	store.Options(sessions.Options{
		MaxAge:   86400,
		Path:     "/",
		Secure:   config.Server.Mode != "debug",
		HttpOnly: true,
	})

	SessionStore = store

}
