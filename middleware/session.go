package middleware

import (
	"memo-server/utils"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionMiddleware() func(c *gin.Context) {
	return sessions.Sessions(
		"session",
		utils.SessionStore,
	)
}
