package controller

import (
	"net/http"

	notifServ "github.com/Group10CapstoneProject/Golang/internal/notifications/service"
	"github.com/labstack/echo/v4"
)

type notificationControllerImpl struct {
	notificationService notifServ.NotificationService
}

// GetNotifications implements NotificationController
func (d *notificationControllerImpl) GetNotifications(c echo.Context) error {
	notifications, err := d.notificationService.PullNotifications(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get notifications",
		"data":    notifications,
	})
}

func NewNotificationController(notificationService notifServ.NotificationService) NotificationController {
	return &notificationControllerImpl{
		notificationService: notificationService,
	}
}
