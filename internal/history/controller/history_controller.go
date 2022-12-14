package controller

import "github.com/labstack/echo/v4"

type HistoryController interface {
	FindHistoryActivity(c echo.Context) error
	FindHistoryOrder(c echo.Context) error
}
