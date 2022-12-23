package controller

import "github.com/labstack/echo/v4"

type UserController interface {
	Signup(c echo.Context) error
	NewAadmin(c echo.Context) error
	GetUser(c echo.Context) error
	GetUsers(c echo.Context) error
	GetAdmins(c echo.Context) error
	UpdateAdmin(c echo.Context) error
	DeleteAdmin(c echo.Context) error
	GetAdminDetail(c echo.Context) error
}
