package controller

import (
	"net/http"
	"strconv"

	"github.com/Group10CapstoneProject/Golang/constans"
	authServ "github.com/Group10CapstoneProject/Golang/internal/auth/service"
	"github.com/Group10CapstoneProject/Golang/internal/paymentMethods/dto"
	paymentMethodServ "github.com/Group10CapstoneProject/Golang/internal/paymentMethods/service"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"github.com/labstack/echo/v4"
)

type paymentMehtodControllerImpl struct {
	paymentMethodService paymentMethodServ.PaymentMethodService
	authService          authServ.AuthService
}

// CreatePaymentMethod implements PaymentMethodController
func (d *paymentMehtodControllerImpl) CreatePaymentMethod(c echo.Context) error {
	claims := d.authService.GetClaims(&c)
	if err := d.authService.ValidationRole(claims, constans.Role_admin, c.Request().Context()); err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
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
	claims := d.authService.GetClaims(&c)
	if err := d.authService.ValidationRole(claims, constans.Role_admin, c.Request().Context()); err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
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
	claims := d.authService.GetClaims(&c)
	if err := d.authService.ValidationToken(claims, c.Request().Context()); err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
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
	claims := d.authService.GetClaims(&c)
	if err := d.authService.ValidationToken(claims, c.Request().Context()); err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	paymentMethods, err := d.paymentMethodService.FindPaymentMethods(c.Request().Context())
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
	claims := d.authService.GetClaims(&c)
	if err := d.authService.ValidationRole(claims, constans.Role_admin, c.Request().Context()); err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
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

func NewPaymentMethodController(paymentMethodService paymentMethodServ.PaymentMethodService, authService authServ.AuthService) PaymentMethodController {
	return &paymentMehtodControllerImpl{
		paymentMethodService: paymentMethodService,
		authService:          authService,
	}
}
