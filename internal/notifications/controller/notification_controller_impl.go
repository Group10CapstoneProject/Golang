package controller

import (
	"net/http"

	"github.com/Group10CapstoneProject/Golang/constans"
	authServ "github.com/Group10CapstoneProject/Golang/internal/auth/service"
	notifServ "github.com/Group10CapstoneProject/Golang/internal/notifications/service"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"github.com/labstack/echo/v4"
)

type notificationControllerImpl struct {
	notificationService notifServ.NotificationService
	authService         authServ.AuthService
}

// GetNotifications implements NotificationController
func (d *notificationControllerImpl) GetNotifications(c echo.Context) error {
	claims := d.authService.GetClaims(&c)
	if err := d.authService.ValidationRole(claims, constans.Role_admin, c.Request().Context()); err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	notifications, err := d.notificationService.PullNotifications(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get notifications",
		"data":    notifications,
	})
}

func NewNotificationController(notificationService notifServ.NotificationService, authService authServ.AuthService) NotificationController {
	return &notificationControllerImpl{
		notificationService: notificationService,
		authService:         authService,
	}
}
