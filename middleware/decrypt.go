package middleware

import (
	"fmt"
	"memo-server/utils"
	"memo-server/utils/r"

	"github.com/gin-gonic/gin"
)

func DecryptMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		isEncrypted := c.Request.Header.Get("Encrypted")
		if isEncrypted != "true" {
			c.Next()
			return
		}

		pem, err := utils.GetCachedPrivateKeyPEM(c)
		if err != nil {
			r.OkJsonResult(c, r.APP_GETRSA_FAILED, "", nil)
			c.Abort()
			return
		}

		fmt.Println(pem)
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

	}

}
