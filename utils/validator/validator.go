package validator

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.Validator.Struct(i)
	if err != nil {
		validationErr := err.(validator.ValidationErrors)
		for _, each := range validationErr {
			switch each.Tag() {
			case "required":
				msg := fmt.Sprintf("%s is required", each.Field())
				return echo.NewHTTPError(http.StatusBadRequest, msg)
			case "len":
				msg := fmt.Sprintf("%s must be %s characters", each.Field(), each.Param())
				return echo.NewHTTPError(http.StatusBadRequest, msg)
			case "email":
				msg := fmt.Sprintf("%s must be email format", each.Field())
				return echo.NewHTTPError(http.StatusBadRequest, msg)
			case "gte":
				msg := fmt.Sprintf("%s must be greater than or equal to %s", each.Field(), each.Param())
				return echo.NewHTTPError(http.StatusBadRequest, msg)
			case "min":
				msg := fmt.Sprintf("%s must be minimum %s characters", each.Field(), each.Param())
				return echo.NewHTTPError(http.StatusBadRequest, msg)
			case "personname":
				msg := fmt.Sprintf("%s not a person name", each.Field())
				return echo.NewHTTPError(http.StatusBadRequest, msg)
			default:
				msg := fmt.Sprintf("Invalid field %s", each.Field())
				return echo.NewHTTPError(http.StatusBadRequest, msg)
			}
		}
	}

	return nil
}

func NewCustomValidator(e *echo.Echo) {
	validator := validator.New()

	// register the custom validator
	if err := validator.RegisterValidation("personname", personNameValidator); err != nil {
		panic(err)
	}

	e.Validator = &CustomValidator{validator}
}

// write custom validator here

func personNameValidator(fl validator.FieldLevel) bool {
	nameRegex := regexp.MustCompile("^[a-zA-Z]+(([',. -][a-zA-Z ])?[a-zA-Z]*)*$")
	return nameRegex.MatchString(fl.Field().String())
}
