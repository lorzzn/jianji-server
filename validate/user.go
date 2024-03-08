package validate

import (
	"jianji-server/model/request"
	"jianji-server/utils"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type User struct{}

func (*User) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		params, _ := utils.GetRequestParams[request.Login](c)
		if err := StructValidate(c, &params,
			validation.Field(&params.Email, validation.Required, validation.Match(EmailRegexp()).Error("请输入正确的邮箱地址")),
			validation.Field(&params.Password, validation.Required),
		); err != nil {
			c.Abort()
			return
		}
		c.Next()
	}
}

func (user *User) Signup() gin.HandlerFunc {
	return user.Login()
}

func (*User) Active() gin.HandlerFunc {
	return func(c *gin.Context) {
		params, _ := utils.GetRequestParams[request.Active](c)
		if err := StructValidate(c, &params,
			validation.Field(&params.State, validation.Required),
		); err != nil {
			c.Abort()
			return
		}
		c.Next()
	}
}

func (*User) RefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		params, _ := utils.GetRequestParams[request.RefreshToken](c)
		if err := StructValidate(c, &params,
			validation.Field(&params.Token, validation.Required),
			validation.Field(&params.RefreshToken, validation.Required),
		); err != nil {
			c.Abort()
			return
		}
		c.Next()
	}
}

func (*User) EditProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		params, _ := utils.GetRequestParams[request.EditProfile](c)
		if err := StructValidate(c, &params,
			validation.Field(&params.Name, validation.Required, validation.RuneLength(1, 16)),
		); err != nil {
			c.Abort()
			return
		}
		c.Next()
	}
}
