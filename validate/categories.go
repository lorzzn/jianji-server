package validate

import (
	"errors"
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
			validation.Field(&params.Data, validation.Each(validation.By(func(value any) error {
				value, ok := value.(request.CreateCategoriesDatum)
				if !ok {
					return errors.New("参数格式不正确")
				}
				return nil
			}))),
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
			validation.Field(&params.Value, validation.By(func(value interface{}) error {
				value, ok := value.([]uint64)
				if !ok {
					return errors.New("参数不正确")
				}
				return nil
			})),
		); err != nil {
			c.Abort()
			return
		}
		c.Next()
	}
}

func (*Categories) CategoryStatistics() gin.HandlerFunc {
	return func(c *gin.Context) {
		params, _ := utils.GetRequestParams[request.CategoryStatistics](c)
		if err := StructValidate(c, &params,
			validation.Field(&params.Value, validation.Required),
		); err != nil {
			c.Abort()
			return
		}
		c.Next()
	}
}
