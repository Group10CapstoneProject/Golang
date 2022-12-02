package controller

import "github.com/labstack/echo/v4"

type NotificationController interface {
	GetNotifications(c echo.Context) error
}
