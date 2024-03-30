package validate

import (
	"jianji-server/model/request"
	"jianji-server/utils"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Categories struct {
}

func (*Categories) CreateCategories() gin.HandlerFunc {
	return func(c *gin.Context) {
		params, _ := utils.GetRequestParams[request.CreateCategories](c)
		if err := StructValidate(c, &params,
			validation.Field(&params.Label, validation.Required),
			validation.Field(&params.ParentValue, validation.Min(uint64(0)).Exclusive()),
		); err != nil {
			c.Abort()
			return
		}
		c.Next()
	}
}

func (*Categories) DeleteCategories() gin.HandlerFunc {
	return func(c *gin.Context) {
		params, _ := utils.GetRequestParams[request.DeleteCategories](c)
		if err := StructValidate(c, &params,
			validation.Field(&params.Value, validation.Required),
		); err != nil {
			c.Abort()
			return
		}
		c.Next()
	}
}
