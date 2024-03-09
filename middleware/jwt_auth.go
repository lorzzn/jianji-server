package middleware

import (
	"jianji-server/utils"
	"jianji-server/utils/r"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		//如果请求没有Authorization, 游客状态
		if authorization == "" {
			c.Next()
			return
		}
		parts := strings.SplitN(authorization, " ", 2)

		//验证token格式
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			r.OkJsonResult(c, r.TOKEN_IS_BAD, "", nil)
			//阻止调用后续的函数
			c.Abort()
			return
		}

		//获取jwt信息
		mc, tokenErr := utils.ParseToken(parts[1])
		if tokenErr != nil {
			//过期token软删除
			if utils.IsTokenValidationErrorExpired(tokenErr) {
				utils.SoftDeleteUserToken(mc.TokenUUID)
			}

			r.OkJsonResult(c, r.TOKEN_AUTHORIZATION_INVALID, "", nil)
			c.Abort()
			return
		}

		//检查jwt是否已拉黑（退出登录...）
		blacklisted, err2 := utils.IsTokenBlacklisted(mc.TokenUUID.String())
		if err2 != nil {
			r.OkJsonResult(c, r.FAIL, "", nil)
			c.Abort()
			return
		}

		if blacklisted {
			r.OkJsonResult(c, r.USER_NOT_LOGIN, "", nil)
			c.Abort()
			return
		}

		// 储存 jwt 信息
		c.Set("UserId", strconv.FormatUint(mc.UserId, 10))
		c.Set("UserUUID", mc.UserUUID.String())
		c.Set("TokenUUID", mc.TokenUUID.String())
		c.Set("Token", parts[1])

		//后续的处理函数可以通过 c.Get("..") 来获取
		c.Next()
	}

}
