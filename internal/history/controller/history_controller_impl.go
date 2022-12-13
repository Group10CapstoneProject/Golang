package controller

import (
	"net/http"

	"github.com/Group10CapstoneProject/Golang/internal/history/dto"
	historyServ "github.com/Group10CapstoneProject/Golang/internal/history/service"
	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/Group10CapstoneProject/Golang/utils/jwt"
	"github.com/labstack/echo/v4"
)

type historyControllerImpl struct {
	historyService historyServ.HistoryService
	jwtService     jwt.JWTService
}

// FindHistoryActivity implements HistoryController
func (d *historyControllerImpl) FindHistoryActivity(c echo.Context) error {
	claims := d.jwtService.GetClaims(&c)
	userId := claims["user_id"].(float64)
	status := c.QueryParam("status")
	qtype := c.QueryParam("type")
	cond := dto.HistoryActivityRequest{
		UserID: uint(userId),
		Status: model.StatusType(status),
		Type:   qtype,
	}
	if err := c.Validate(cond); err != nil {
		return err
	}
	history, err := d.historyService.FindHistoryActivity(&cond, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get history activity",
		"data":    history,
	})
}

// FindHistoryOrder implements HistoryController
func (d *historyControllerImpl) FindHistoryOrder(c echo.Context) error {
	claims := d.jwtService.GetClaims(&c)
	userId := claims["user_id"].(float64)
	qtype := c.QueryParam("type")
	cond := dto.HistoryOrderRequest{
		UserID: uint(userId),
		Type:   qtype,
	}
	if err := c.Validate(cond); err != nil {
		return err
	}
	history, err := d.historyService.FindHistoryOrder(&cond, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get history order",
		"data":    history,
	})
}

func NewHistoryController(historyService historyServ.HistoryService, jwtService jwt.JWTService) HistoryController {
	return &historyControllerImpl{
		historyService: historyService,
		jwtService:     jwtService,
	}
}
