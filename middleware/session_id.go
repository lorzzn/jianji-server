package middleware

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SessionIdMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionId := session.Get("SessionId")
		if sessionId == nil {
			sessionId = uuid.New().String()
			session.Set("SessionId", sessionId)
			session.Save()
		}
		c.Set("SessionId", sessionId)
		c.Next()
	}
}
