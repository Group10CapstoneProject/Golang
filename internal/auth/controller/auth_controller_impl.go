package controller

import (
	"net/http"

	"github.com/Group10CapstoneProject/Golang/internal/auth/dto"
	authService "github.com/Group10CapstoneProject/Golang/internal/auth/service"
	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/labstack/echo/v4"
)

type authControllerImpl struct {
	authService authService.AuthService
}

// InitRoute implements AuthController
func (d *authControllerImpl) InitRoute(api *echo.Group, protect *echo.Group) {
	auth := api.Group("/auth")

	auth.POST("/login", d.Login)
	auth.POST("/refresh", d.RefreshToken)

	protect.POST("/auth/logout", d.Logout)
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
	if err := d.authService.ValidationToken(claims, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
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

func NewAuthController(auth authService.AuthService) AuthController {
	return &authControllerImpl{
		authService: auth,
	}
}