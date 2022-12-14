package controller

import "github.com/labstack/echo/v4"

type FileController interface {
	Upload(c echo.Context) error
}
