package controller

import (
	"net/http"
	"strconv"

	"github.com/Group10CapstoneProject/Golang/internal/payment_methods/dto"
	paymentMethodServ "github.com/Group10CapstoneProject/Golang/internal/payment_methods/service"
	jwtServ "github.com/Group10CapstoneProject/Golang/utils/jwt"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"github.com/labstack/echo/v4"
)

type paymentMehtodControllerImpl struct {
	paymentMethodService paymentMethodServ.PaymentMethodService
	jwtService           jwtServ.JWTService
}

// CreatePaymentMethod implements PaymentMethodController
func (d *paymentMehtodControllerImpl) CreatePaymentMethod(c echo.Context) error {
	var paymentMethod dto.PaymentMethodStoreRequest
	if err := c.Bind(&paymentMethod); err != nil {
		return err
	}
	if err := c.Validate(paymentMethod); err != nil {
		return err
	}
	if err := d.paymentMethodService.CreatePaymentMethod(&paymentMethod, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "new payment method succcess created",
	})
}

// DeletePaymentMethod implements PaymentMethodController
func (d *paymentMehtodControllerImpl) DeletePaymentMethod(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "record not found")
	}
	if err := d.paymentMethodService.DeletePaymentMethod(uint(id), c.Request().Context()); err != nil {
		if err == myerrors.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete payment method",
	})
}

// GetPaymentMethodDetail implements PaymentMethodController
func (d *paymentMehtodControllerImpl) GetPaymentMethodDetail(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "record not found")
	}
	paymentMethod, err := d.paymentMethodService.FindPaymentMethodById(uint(id), c.Request().Context())
	if err != nil {
		if err == myerrors.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get payment method",
		"data":    paymentMethod,
	})
}

// GetPaymentMethods implements PaymentMethodController
func (d *paymentMehtodControllerImpl) GetPaymentMethods(c echo.Context) error {
	qParam := c.QueryParam("access")
	access := false
	if qParam != "" {
		temp, err := strconv.ParseBool(qParam)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		access = temp
	}
	paymentMethods, err := d.paymentMethodService.FindPaymentMethods(access, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get payment methods",
		"data":    paymentMethods,
	})
}

// UpdatePaymentMethod implements PaymentMethodController
func (d *paymentMehtodControllerImpl) UpdatePaymentMethod(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "record not found")
	}
	var paymentMethod dto.PaymentMethodUpdateRequest
	if err := c.Bind(&paymentMethod); err != nil {
		return err
	}
	if err := c.Validate(paymentMethod); err != nil {
		return err
	}
	paymentMethod.ID = uint(id)
	if err := d.paymentMethodService.UpdatePaymentMethod(&paymentMethod, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update payment method",
	})
}

func NewPaymentMethodController(paymentMethodService paymentMethodServ.PaymentMethodService, jwtService jwtServ.JWTService) PaymentMethodController {
	return &paymentMehtodControllerImpl{
		paymentMethodService: paymentMethodService,
		jwtService:           jwtService,
	}
}
