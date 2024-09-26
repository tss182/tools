package tools

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
)

func HandlerBindingError(c *gin.Context, obj any, shouldType string) error {
	var err error
	switch shouldType {
	case ShouldTypeQuery:
		err = c.ShouldBindQuery(obj)
	case ShouldTypeJson:
		err = c.ShouldBindJSON(obj)
	case ShouldTypeForm:
		err = c.ShouldBind(obj)
	}
	return ErrorHandle(err)
}

func HandlerBindingErrorFiber(c *fiber.Ctx, obj any, shouldType string) error {
	var err error
	switch shouldType {
	case ShouldTypeQuery:
		err = c.BodyParser(obj)
	case ShouldTypeJson:
		err = c.BodyParser(obj)
	case ShouldTypeForm:
		err = c.BodyParser(obj)
	}
	return ErrorHandle(err)
}

func ErrorHandle(err error) error {
	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		var required []string
		var errorSlice []string
		for _, f := range errs {
			switch f.Tag() {
			case "required":
				required = append(required, f.Field())
			case "min":
				i, _ := strconv.Atoi(f.Param())
				str := fmt.Sprintf("minimum %s is %s", f.Field(), DecimalSeparator(i))
				errorSlice = append(errorSlice, str)
			case "max":
				i, _ := strconv.Atoi(f.Param())
				str := fmt.Sprintf("maximum %s is %s", f.Field(), DecimalSeparator(i))
				errorSlice = append(errorSlice, str)
			case "stringNumberOnly":
				str := fmt.Sprintf("%s just support letter and number", f.Field())
				errorSlice = append(errorSlice, str)
			default:
				str := fmt.Sprintf("error %s is %s", f.Field(), f.Tag())
				errorSlice = append(errorSlice, str)
			}
		}
		var errResult []string
		if len(required) > 0 {
			errResult = append(errResult, strings.Join(required, ", ")+" is required")
		}
		errResult = append(errResult, errorSlice...)
		return errors.New(strings.Join(errResult, ", "))
	}
	return err
}
