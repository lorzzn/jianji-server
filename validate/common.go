package validate

import (
	"fmt"
	"memo-server/utils/r"
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func StructValidate[T any](c *gin.Context, data *T, rules ...*validation.FieldRules) error {
	err := validation.ValidateStruct(data, rules...)
	if err != nil {
		r.Result(c, http.StatusOK, r.ERROR_INVALID_PARAM, fmt.Sprint(err), nil)
	}
	return err
}
