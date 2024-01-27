package validate

import (
	"memo-server/model/request"
	"memo-server/utils"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct{}

func (*User) Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		params, _ := utils.GetRequestParams[request.Signup](c)
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
