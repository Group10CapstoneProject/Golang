package controller

import "github.com/labstack/echo/v4"

type OfflineClassController interface {
	InitRoute(api *echo.Group)
	// Offline Class
	CreateOfflineClass(c echo.Context) error
	GetOfflineClass(c echo.Context) error
	GetOfflineClassDetail(c echo.Context) error
	UpdateOfflineClass(c echo.Context) error
	DeleteOfflineClass(c echo.Context) error
}
