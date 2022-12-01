package controller

import "github.com/labstack/echo/v4"

type FileController interface {
	InitRoute(api *echo.Group)
	Upload(c echo.Context) error
}
