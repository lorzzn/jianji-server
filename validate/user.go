package validate

import (
	"jianji-server/model/request"
	"jianji-server/utils"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct{}

func (*User) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		params, _ := utils.GetRequestParams[request.Login](c)
		if err := StructValidate(c, &params,
			validation.Field(&params.Email, validation.Required, is.Email),
			validation.Field(&params.Password, validation.Required),
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
