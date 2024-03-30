package v1

import (
	"jianji-server/utils/r"

	"github.com/gin-gonic/gin"
)

type User struct{}

func (*User) Login(c *gin.Context) {
	var code, message, data = UserService.Login(c)
	//如果用户不存在就进行注册, 后面交给下一个handler Signup，否则返回登陆结果并中止
	if code == r.USER_NOT_EXISTED {
		c.Next()
	} else {
		r.OkJsonResult(c, code, message, data)
		c.Abort()
	}
}

func (*User) Signup(c *gin.Context) {
	var code, message, data = UserService.Signup(c)
	r.OkJsonResult(c, code, message, data)
}

func (*User) Active(c *gin.Context) {
	var code, message, data = UserService.Active(c)
	//如果有报错就中止执行
	if code == r.OK {
		c.Next()
	} else {
		r.OkJsonResult(c, code, message, data)
		c.Abort()
	}
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
