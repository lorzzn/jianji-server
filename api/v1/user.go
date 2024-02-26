package v1

import (
	"jianji-server/utils/r"

	"github.com/gin-gonic/gin"
)

type User struct{}

func (*User) Login(c *gin.Context) {
	var code, message, data = UserService.Login(c)
	r.OkJsonResult(c, code, message, data)
}

func (*User) RefreshToken(c *gin.Context) {
	var code, message, data = UserService.RefreshToken(c)
	r.OkJsonResult(c, code, message, data)
}
