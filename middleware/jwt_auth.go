package middleware

import (
	"memo-server/utils"
	"memo-server/utils/r"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		//如果请求没有Authorization，就在响应中添加
		if authorization == "" {
			accessToken, refreshToken, _ := utils.GenToken(false, nil)
			authorization = "Bearer " + accessToken
			c.Header("Authorization", authorization)
			c.Header("Refresh-Token", refreshToken)
		}
		parts := strings.SplitN(authorization, " ", 2)

		if !(len(parts) == 2 && parts[0] == "Bearer") {
			r.OkJsonResult(c, r.JWT_BAD_AUTHORIZATION, "", nil)
			//阻止调用后续的函数
			c.Abort()
			return
		}
		mc, err := utils.ParseToken(parts[1])
		if err != nil {
			r.OkJsonResult(c, r.JWT_AUTHORIZATION_INVALID, "", nil)
			c.Abort()
			return
		}

		// 储存 jwt 信息
		c.Set("IsLoggedIn", mc.IsLoggedIn)
		c.Set("UserId", mc.UserId)
		c.Set("JwtId", mc.JwtId)

		//后续的处理函数可以通过 c.Get("..") 来获取
		c.Next()
	}

}
