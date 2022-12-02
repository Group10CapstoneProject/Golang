package controller

import (
	"github.com/labstack/echo/v4"
)

type PaymentMethodController interface {
	CreatePaymentMethod(c echo.Context) error
	GetPaymentMethods(c echo.Context) error
	GetPaymentMethodDetail(c echo.Context) error
	UpdatePaymentMethod(c echo.Context) error
	DeletePaymentMethod(c echo.Context) error
}
