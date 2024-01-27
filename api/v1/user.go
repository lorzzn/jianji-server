package v1

import (
	"memo-server/utils/r"

	"github.com/gin-gonic/gin"
)

type User struct{}

func (*User) Signup(c *gin.Context) {
	var code, message, data = UserService.Signup(c)
	r.OkResult(c, code, message, data)
}
