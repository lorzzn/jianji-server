package middleware

import (
	"jianji-server/utils"

	"github.com/gin-gonic/gin"
)

func TxQueryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tx := utils.DB.Begin()
		c.Set("TxQuery", tx)

		c.Next()

		if c.IsAborted() {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}
}
