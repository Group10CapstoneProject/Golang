package controller

import "github.com/labstack/echo/v4"

type UserController interface {
	Signup(c echo.Context) error
	NewAadmin(c echo.Context) error
	GetUser(c echo.Context) error
	GetUsers(c echo.Context) error
	GetAdmins(c echo.Context) error
	InitRoute(api *echo.Group, protect *echo.Group)
}
