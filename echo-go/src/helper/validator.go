package helper

import (
	"fmt"
	"net/http"

	. "echo-go/src/log"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func Validate(r *echo.Echo) {
	r.Validator = &CustomValidator{validator: validator.New()}
	r.HTTPErrorHandler = func(err error, ctx echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required":
					report.Message = fmt.Sprintf("%s is required.", err.Field())
				case "email":
					report.Message = fmt.Sprintf("%s is not valid email.", err.Field())
				case "gte":
					report.Message = fmt.Sprintf("%s value must be greater than %s", err.Field(), err.Param())
				case "lte":
					report.Message = fmt.Sprintf("%s value must be lower than %s", err.Field(), err.Param())
				}
				break
			}
		}

		MakeLogEntry(ctx).Error(report.Message)
		// ctx.Logger().Error(report)
		ctx.JSON(report.Code, report)
	}
}
