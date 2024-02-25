package validate

import (
	"fmt"
	"jianji-server/utils/r"
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func StructValidate[T any](c *gin.Context, data *T, rules ...*validation.FieldRules) error {
	err := validation.ValidateStruct(data, rules...)
	if err != nil {
		r.JsonResult(c, http.StatusOK, r.ERROR_INVALID_PARAM, fmt.Sprint(err), nil)
	}
	return err
}
