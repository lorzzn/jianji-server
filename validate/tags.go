package validate

import (
	"errors"
	"jianji-server/model/request"
	"jianji-server/utils"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Tags struct {
}

func (*Tags) CreateTags() gin.HandlerFunc {
	return func(c *gin.Context) {
		params, _ := utils.GetRequestParams[request.CreateTags](c)
		if err := StructValidate(c, &params,
			validation.Field(&params.Data, validation.Each(validation.By(func(value any) error {
				value, ok := value.(request.CreateTagsDatum)
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

func (*Tags) DeleteTags() gin.HandlerFunc {
	return func(c *gin.Context) {
		params, _ := utils.GetRequestParams[request.DeleteTags](c)
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

func (*Tags) TagStatistics() gin.HandlerFunc {
	return func(c *gin.Context) {
		params, _ := utils.GetRequestParams[request.TagStatistics](c)
		if err := StructValidate(c, &params,
			validation.Field(&params.Value, validation.Required),
		); err != nil {
			c.Abort()
			return
		}
		c.Next()
	}
}
