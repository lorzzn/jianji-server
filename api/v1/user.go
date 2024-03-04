package v1

import (
	"jianji-server/utils/r"

	"github.com/gin-gonic/gin"
)

type User struct{}

func (*User) Login(c *gin.Context) {
	var code, message, data = UserService.Login(c)
	needSignup, exists := c.Get("NeedSignup")
	//如果不需要注册，返回结果后终止
	if !exists || !needSignup.(bool) {
		r.OkJsonResult(c, code, message, data)
		c.Abort()
		return
	}
}

func (*User) Signup(c *gin.Context) {
	var code, message, data = UserService.Signup(c)
	r.OkJsonResult(c, code, message, data)
}

func (*User) Active(c *gin.Context) {
	var code, message, data = UserService.Active(c)
	r.OkJsonResult(c, code, message, data)
}

func (*User) Logout(c *gin.Context) {
	var code, message = UserService.Logout(c)
	r.OkJsonResult(c, code, message, nil)
}

func (*User) RefreshToken(c *gin.Context) {
	var code, message, data = UserService.RefreshToken(c)
	r.OkJsonResult(c, code, message, data)
}

func (*User) GetProfile(c *gin.Context) {
	var code, message, data = UserService.GetProfile(c)
	r.OkJsonResult(c, code, message, data)
}

func (*User) EditProfile(c *gin.Context) {
	var code, message, data = UserService.EditProfile(c)
	r.OkJsonResult(c, code, message, data)
}
