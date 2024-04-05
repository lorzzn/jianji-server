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
		params, _ := utils.GetRequestParams[request.CreateTag](c)
		if err := StructValidate(c, &params,
			validation.Field(&params.Data, validation.Each(validation.By(func(value any) error {
				value, ok := value.(request.CreateTagDatum)
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
		params, _ := utils.GetRequestParams[request.DeleteTagBatch](c)
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
