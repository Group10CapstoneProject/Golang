package controller

import (
	"net/http"

	dashboardServ "github.com/Group10CapstoneProject/Golang/internal/dashboard/service"
	"github.com/labstack/echo/v4"
)

type dashboardControllerImpl struct {
	dashboardService dashboardServ.DashboardService
}

// GetDashboard implements DashboardController
func (d *dashboardControllerImpl) GetDashboard(c echo.Context) error {
	result, err := d.dashboardService.GetDashboard(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get dashboard",
		"data":    result,
	})
}

func NewDashboardController(dashboardService dashboardServ.DashboardService) DashboardController {
	return &dashboardControllerImpl{
		dashboardService: dashboardService,
	}
}
