package v1

import (
	"memo-server/utils/r"

	"github.com/gin-gonic/gin"
)

type User struct{}

func (*User) Login(c *gin.Context) {
	var code, message, data = UserService.Login(c)
	r.OkJsonResult(c, code, message, data)
}