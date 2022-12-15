package validator

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

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
			case "name":
				msg := fmt.Sprintf("%s use an invalid character", each.Field())
				return echo.NewHTTPError(http.StatusBadRequest, msg)
			case "alpha":
				msg := fmt.Sprintf("%s must alphabet character", each.Field())
				return echo.NewHTTPError(http.StatusBadRequest, msg)
			case "url":
				msg := fmt.Sprintf("%s not an url", each.Field())
				return echo.NewHTTPError(http.StatusBadRequest, msg)
			case "unique":
				msg := fmt.Sprintf("%s have duplicate data", each.Field())
				return echo.NewHTTPError(http.StatusBadRequest, msg)
			case "number":
				msg := fmt.Sprintf("%s must number character", each.Field())
				return echo.NewHTTPError(http.StatusBadRequest, msg)
			case "status":
				msg := fmt.Sprintf("%s not a status or status not allowed", each.Field())
				return echo.NewHTTPError(http.StatusBadRequest, msg)
			case "image":
				msg := fmt.Sprintf("%s not an image (invalid extention)", each.Field())
				return echo.NewHTTPError(http.StatusBadRequest, msg)
			case "title":
				msg := fmt.Sprintf("%s not allowed", each.Field())
				return echo.NewHTTPError(http.StatusBadRequest, msg)
			case "mytime":
				msg := fmt.Sprintf("%s invalid format time, use (YYYY-MM-DD hh:mm:ss), example 2006-01-02 15:04:05", each.Field())
				return echo.NewHTTPError(http.StatusBadRequest, msg)
			case "mydate":
				msg := fmt.Sprintf("%s invalid format date, use (YYYY-MM-DD), example 2006-01-02", each.Field())
				return echo.NewHTTPError(http.StatusBadRequest, msg)
			case "dob":
				msg := fmt.Sprintf("%s invalid date of birth", each.Field())
				return echo.NewHTTPError(http.StatusBadRequest, msg)
			case "activity":
				msg := fmt.Sprintf("%s invalid or status not allowed", each.Field())
				return echo.NewHTTPError(http.StatusBadRequest, msg)
			case "ordertype":
				msg := fmt.Sprintf("%s must DESC|desc (descending) or ASC|asc (ascending)", each.Field())
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
	if err := validator.RegisterValidation("name", nameValidator); err != nil {
		panic(err)
	}
	if err := validator.RegisterValidation("status", statusValidator); err != nil {
		panic(err)
	}
	if err := validator.RegisterValidation("image", imageValidator); err != nil {
		panic(err)
	}
	if err := validator.RegisterValidation("title", titleValidator); err != nil {
		panic(err)
	}
	if err := validator.RegisterValidation("mytime", mytimeValidator); err != nil {
		panic(err)
	}
	if err := validator.RegisterValidation("activity", activityValidator); err != nil {
		panic(err)
	}
	if err := validator.RegisterValidation("dob", dobValidator); err != nil {
		panic(err)
	}
	if err := validator.RegisterValidation("mydate", mydateValidator); err != nil {
		panic(err)
	}
	if err := validator.RegisterValidation("ordertype", orderTypeValidator); err != nil {
		panic(err)
	}

	e.Validator = &CustomValidator{validator}
}

// write custom validator here

func personNameValidator(fl validator.FieldLevel) bool {
	nameRegex := regexp.MustCompile("^[a-zA-Z]+(([',. -][a-zA-Z ])?[a-zA-Z]*)*$")
	return nameRegex.MatchString(fl.Field().String())
}

func nameValidator(fl validator.FieldLevel) bool {
	nameRegex := regexp.MustCompile("^[a-zA-Z0-9]+(([',. -][a-zA-Z0-9 ])?[a-zA-Z0-9]*)*$")
	return nameRegex.MatchString(fl.Field().String())
}

func statusValidator(fl validator.FieldLevel) bool {
	nameRegex := regexp.MustCompile("^(ACTIVE|REJECT|INACTIVE)*$")
	return nameRegex.MatchString(fl.Field().String())
}

func activityValidator(fl validator.FieldLevel) bool {
	nameRegex := regexp.MustCompile("^(ACTIVE|WAITING)*$")
	return nameRegex.MatchString(fl.Field().String())
}

func titleValidator(fl validator.FieldLevel) bool {
	nameRegex := regexp.MustCompile("^(member_booking|offline_class_booking|online_class_booking|trainer|profile|member_type|online_class|offline_class|online_class_category|offline_class_category)*$")
	return nameRegex.MatchString(fl.Field().String())
}

func imageValidator(fl validator.FieldLevel) bool {
	nameRegex := regexp.MustCompile(`^.*\.(jpg|JPG|png|PNG|jpeg|JPEG)$`)
	return nameRegex.MatchString(fl.Field().String())
}

func mytimeValidator(fl validator.FieldLevel) bool {
	layoutFormat := "2006-01-02 15:04:05"
	_, err := time.Parse(layoutFormat, fl.Field().String())
	return err == nil
}

func dobValidator(fl validator.FieldLevel) bool {
	layoutFormat := "2006-01-02"
	dob, err := time.Parse(layoutFormat, fl.Field().String())
	if err != nil {
		return false
	}
	now := time.Now()
	return dob.Before(now)
}

func mydateValidator(fl validator.FieldLevel) bool {
	layoutFormat := "2006-01-02"
	_, err := time.Parse(layoutFormat, fl.Field().String())
	return err == nil
}

func orderTypeValidator(fl validator.FieldLevel) bool {
	nameRegex := regexp.MustCompile("^(DESC|ASC|desc|asc)*$")
	return nameRegex.MatchString(fl.Field().String())
}
