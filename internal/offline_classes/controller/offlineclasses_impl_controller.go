package controller

import (
	"net/http"
	"strconv"

	"github.com/Group10CapstoneProject/Golang/config"
	"github.com/Group10CapstoneProject/Golang/constans"
	authServ "github.com/Group10CapstoneProject/Golang/internal/auth/service"
	"github.com/Group10CapstoneProject/Golang/internal/members/dto"
	offlineclassServ "github.com/Group10CapstoneProject/Golang/internal/members/service"
	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type offlineclassControllerImpl struct {
	offlineclassService offlineclassServ.OfflineClassService
	authService         authServ.AuthService
}

// InitRoute implements OfflineClassController
func (d *offlineclassControllerImpl) InitRoute(api *echo.Group) {
	offlineclass := api.Group("/offlineclass", middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))

	offlineclass.POST("", d.CreateOfflineClass)
	offlineclass.GET("", d.GetOfflineClass)
	offlineclass.GET("/:id", d.GetOfflineClassDetail)
	offlineclass.PUT("/:id", d.UpdateOfflineClass)
	offlineclass.DELETE("/:id", d.DeleteOfflineClass)

}

// CreateOfflineClass implements OfflineClassController
func (d *offlineclassControllerImpl) CreateOfflineClass(c echo.Context) error {
	claims := d.authService.GetClaims(&c)
	if err := d.authService.ValidationToken(claims, c.Request().Context()); err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	var offlineclass dto.OfflineClassRequest
	if err := c.Bind(&offlineclass); err != nil {
		return err
	}
	if err := c.Validate(offlineclass); err != nil {
		return err
	}
	if err := d.offlineclassService.CreateOfflineClass(&offlineclass, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "new offlineclass success created",
	})
}

// DeleteMember implements OfflineClassController
func (d *offlineclassControllerImpl) DeleteOfflineClass(c echo.Context) error {
	claims := d.authService.GetClaims(&c)
	if err := d.authService.ValidationRole(claims, constans.Role_admin, c.Request().Context()); err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "record not found")
	}
	if err := d.offlineclassService.DeleteOfflineClass(uint(id), c.Request().Context()); err != nil {
		if err == myerrors.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete offlineclass",
	})
}

// GetOfflineClassDetail implements OfflineClassController
func (d *offlineclassControllerImpl) GetOfflineClassDetail(c echo.Context) error {
	claims := d.authService.GetClaims(&c)
	if err := d.authService.ValidationToken(claims, c.Request().Context()); err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "record not found")
	}
	offlineclass, err := d.offlineclassService.FindOfflineClassById(uint(id), c.Request().Context())
	if err != nil {
		if err == myerrors.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get offlineclass",
		"data":    offlineclass,
	})
}

// GetOfflineClasses implements OfflineClassController
func (d *offlineclassControllerImpl) GetOfflineClass(c echo.Context) error {
	claims := d.authService.GetClaims(&c)
	if err := d.authService.ValidationRole(claims, constans.Role_admin, c.Request().Context()); err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	var query model.Pagination
	query.NewPageQuery(c)

	offlineclass, err := d.offlineclassService.FindOfflineClass(&query, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get members",
		"data":    offlineclass,
	})
}

// UpdateOfflineClass implements OfflineClassController
func (d *offlineclassControllerImpl) UpdateOfflineClass(c echo.Context) error {
	claims := d.authService.GetClaims(&c)
	if err := d.authService.ValidationToken(claims, c.Request().Context()); err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "record not found")
	}
	var offlineclass dto.OfflineClassRequest
	if err := c.Bind(&offlineclass); err != nil {
		return err
	}
	if err := c.Validate(offlineclass); err != nil {
		return err
	}
	offlineclass.ID = uint(id)
	if err := d.offlineclassService.UpdateOfflineClass(&offlineclass, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update offlineclass",
	})
}

func NewOfflineClassController(offlineclassService offlineclassServ.OfflineClassService, authService authServ.AuthService) OfflineClassController {
	return &offlineclassControllerImpl{
		offlineclassService: offlineclassService,
		authService:         authService,
	}
}
