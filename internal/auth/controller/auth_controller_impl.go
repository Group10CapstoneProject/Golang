package controller

import (
	"net/http"
	"time"

	"github.com/Group10CapstoneProject/Golang/constans"
	"github.com/Group10CapstoneProject/Golang/internal/auth/dto"
	authService "github.com/Group10CapstoneProject/Golang/internal/auth/service"
	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"github.com/labstack/echo/v4"
)

type authControllerImpl struct {
	authService authService.AuthService
}

// Login implements AuthController
func (d *authControllerImpl) Login(c echo.Context) error {
	var credential dto.UserCredential
	if err := c.Bind(&credential); err != nil {
		return err
	}
	if err := c.Validate(credential); err != nil {
		return err
	}
	token, err := d.authService.Login(credential, c.Request().Context())
	if err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "login success",
		"data":    token,
	})
}

// Logout implements AuthController
func (d *authControllerImpl) Logout(c echo.Context) error {
	claims := d.authService.GetClaims(&c)
	userId := claims["user_id"].(float64)
	if err := d.authService.Logout(uint(userId), c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "logout success",
	})
}

// RefreshToken implements AuthController
func (d *authControllerImpl) RefreshToken(c echo.Context) error {
	var token model.Token
	if err := c.Bind(&token); err != nil {
		return err
	}
	if err := c.Validate(token); err != nil {
		return err
	}
	newToken, err := d.authService.Refresh(token, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "refresh success",
		"data":    newToken,
	})
}

// RefreshAdminToken implements AuthController
func (d *authControllerImpl) RefreshAdminToken(c echo.Context) error {
	var token model.Token
	if err := c.Bind(&token); err != nil {
		return err
	}
	if err := c.Validate(token); err != nil {
		return err
	}
	newToken, err := d.authService.Refresh(token, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	response := model.AdminToken{
		Access: model.Access{
			AccessAccessToken: newToken.AccessToken,
			ExpiredAt:         time.Now().Add(constans.ExpAccessToken),
		},
		Refresh: model.Refresh{
			RefreshToken: newToken.RefreshToken,
			ExpiredAt:    time.Now().Add(constans.ExpRefreshToken),
		},
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "refresh success",
		"data":    response,
	})
}

// LoginAdmin implements AuthController
func (d *authControllerImpl) LoginAdmin(c echo.Context) error {
	var credential dto.UserCredential
	if err := c.Bind(&credential); err != nil {
		return err
	}
	if err := c.Validate(credential); err != nil {
		return err
	}
	token, err := d.authService.LoginAdmin(credential, c.Request().Context())
	if err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "login success",
		"data":    token,
	})
}

func NewAuthController(auth authService.AuthService) AuthController {
	return &authControllerImpl{
		authService: auth,
	}
}
