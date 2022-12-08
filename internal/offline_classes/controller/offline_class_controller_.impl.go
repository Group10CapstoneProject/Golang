package controller

import (
	"net/http"
	"strconv"

	"github.com/Group10CapstoneProject/Golang/constans"
	authServ "github.com/Group10CapstoneProject/Golang/internal/auth/service"
	memberServ "github.com/Group10CapstoneProject/Golang/internal/members/service"
	dtoNotif "github.com/Group10CapstoneProject/Golang/internal/notifications/dto"
	notificationServ "github.com/Group10CapstoneProject/Golang/internal/notifications/service"
	"github.com/Group10CapstoneProject/Golang/internal/offline_classes/dto"
	offlineClassServ "github.com/Group10CapstoneProject/Golang/internal/offline_classes/service"
	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type offlineclassControllerImpl struct {
	offlineClassService offlineClassServ.OfflineClassService
	memberService       memberServ.MemberService
	authService         authServ.AuthService
	notificationService notificationServ.NotificationService
}

// CheckOfflineClassBooking implements OfflineClassController
func (d *offlineclassControllerImpl) CheckOfflineClassBooking(c echo.Context) error {
	claims := d.authService.GetClaims(&c)
	if err := d.authService.ValidationRole(claims, constans.Role_admin, c.Request().Context()); err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	emailParam := c.QueryParam("email")
	codeParam := c.QueryParam("code")
	code, err := uuid.Parse(codeParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	condition := dto.TakeOfflineClassBooking{
		Email: emailParam,
		Code:  code,
	}
	if err := c.Validate(condition); err != nil {
		return err
	}
	offlineClassBooking, err := d.offlineClassService.CheckOfflineClassBookings(&condition, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get offline class booking",
		"data":    offlineClassBooking,
	})
}

// CreateOfflineClass implements OfflineClassController
func (d *offlineclassControllerImpl) CreateOfflineClass(c echo.Context) error {
	claims := d.authService.GetClaims(&c)
	if err := d.authService.ValidationRole(claims, constans.Role_admin, c.Request().Context()); err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	var offlineClass dto.OfflineClassStoreRequest
	if err := c.Bind(&offlineClass); err != nil {
		return err
	}
	if err := c.Validate(offlineClass); err != nil {
		return err
	}
	if err := d.offlineClassService.CreateOfflineClass(&offlineClass, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "new offline class success created",
	})
}

// CreateOfflineClassBooking implements OfflineClassController
func (d *offlineclassControllerImpl) CreateOfflineClassBooking(c echo.Context) error {
	claims := d.authService.GetClaims(&c)
	if err := d.authService.ValidationToken(claims, c.Request().Context()); err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	var offlineClassBooking dto.OfflineClassBookingStoreRequest
	if err := c.Bind(&offlineClassBooking); err != nil {
		return err
	}
	if err := c.Validate(offlineClassBooking); err != nil {
		return err
	}
	offlineClassBooking.UserID = uint(claims["user_id"].(float64))
	if err := d.offlineClassService.CreateOfflineClassBooking(&offlineClassBooking, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "new offline class booking success created",
	})
}

// CreateOfflineClassCategory implements OfflineClassController
func (d *offlineclassControllerImpl) CreateOfflineClassCategory(c echo.Context) error {
	claims := d.authService.GetClaims(&c)
	if err := d.authService.ValidationRole(claims, constans.Role_admin, c.Request().Context()); err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	var offlineClassCategory dto.OfflineClassCategoryStoreRequest
	if err := c.Bind(&offlineClassCategory); err != nil {
		return err
	}
	if err := c.Validate(offlineClassCategory); err != nil {
		return err
	}
	if err := d.offlineClassService.CreateOfflineClassCategory(&offlineClassCategory, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "new offline class category success created",
	})
}

// DeleteOfflineClass implements OfflineClassController
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
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	if err := d.offlineClassService.DeleteOfflineClass(uint(id), c.Request().Context()); err != nil {
		if err == myerrors.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete offline class",
	})
}

// DeleteOfflineClassBooking implements OfflineClassController
func (d *offlineclassControllerImpl) DeleteOfflineClassBooking(c echo.Context) error {
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
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	if err := d.offlineClassService.DeleteOfflineClassBooking(uint(id), c.Request().Context()); err != nil {
		if err == myerrors.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete offline class booking",
	})
}

// DeleteOfflineClassCategory implements OfflineClassController
func (d *offlineclassControllerImpl) DeleteOfflineClassCategory(c echo.Context) error {
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
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	if err := d.offlineClassService.DeleteOfflineClassCategory(uint(id), c.Request().Context()); err != nil {
		if err == myerrors.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete offline class category",
	})
}

// GetOfflineClassBookingDetail implements OfflineClassController
func (d *offlineclassControllerImpl) GetOfflineClassBookingDetail(c echo.Context) error {
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
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	offlineClassBooking, err := d.offlineClassService.FindOfflineClassBookingById(uint(id), c.Request().Context())
	if err != nil {
		if err == myerrors.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if claims["role"].(string) == constans.Role_superadmin || claims["role"].(string) == constans.Role_admin {
		notif := dtoNotif.NotificationReadRequest{
			TransactionID: uint(id),
			Title:         "Offline Class",
		}
		if err := d.notificationService.ReadNotification(&notif, c.Request().Context()); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get offline class booking",
		"data":    offlineClassBooking,
	})
}

// GetOfflineClassBookingUser implements OfflineClassController
func (d *offlineclassControllerImpl) GetOfflineClassBookingUser(c echo.Context) error {
	claims := d.authService.GetClaims(&c)
	if err := d.authService.ValidationRole(claims, constans.Role_user, c.Request().Context()); err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	offlineClassBooking, err := d.offlineClassService.FindOfflineClassBookingByUser(uint(claims["user_id"].(float64)), c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get offline class booking",
		"data":    offlineClassBooking,
	})
}

// GetOfflineClassBookings implements OfflineClassController
func (d *offlineclassControllerImpl) GetOfflineClassBookings(c echo.Context) error {
	claims := d.authService.GetClaims(&c)
	if err := d.authService.ValidationRole(claims, constans.Role_admin, c.Request().Context()); err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	var query model.Pagination
	query.NewPageQuery(c)

	offlineClassBookings, err := d.offlineClassService.FindOfflineClassBookings(&query, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get offline class bookings",
		"data":    offlineClassBookings,
	})
}

// GetOfflineClassCategories implements OfflineClassController
func (d *offlineclassControllerImpl) GetOfflineClassCategories(c echo.Context) error {
	claims := d.authService.GetClaims(&c)
	if err := d.authService.ValidationToken(claims, c.Request().Context()); err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	offlineClassCategories, err := d.offlineClassService.FindOfflineClassCategories(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get offline class categories",
		"data":    offlineClassCategories,
	})
}

// GetOfflineClassCategoryDetail implements OfflineClassController
func (d *offlineclassControllerImpl) GetOfflineClassCategoryDetail(c echo.Context) error {
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
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	offlineClassCategory, err := d.offlineClassService.FindOfflineClassCategoryById(uint(id), c.Request().Context())
	if err != nil {
		if err == myerrors.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get offline class category",
		"data":    offlineClassCategory,
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
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	offlineClass, err := d.offlineClassService.FindOfflineClassById(uint(id), c.Request().Context())
	if err != nil {
		if err == myerrors.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	offlineClass.AccessClass = true

	if claims["role"].(string) == constans.Role_user {
		memberUser, err := d.memberService.FindMemberByUser(uint(claims["user_id"].(float64)), c.Request().Context())
		if err != nil {
			if err != myerrors.ErrRecordNotFound {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			offlineClass.AccessClass = false
		} else {
			offlineClass.AccessClass = memberUser.MemberType.AccessOfflineClass
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get offline class",
		"data":    offlineClass,
	})
}

// GetOfflineClasses implements OfflineClassController
func (d *offlineclassControllerImpl) GetOfflineClasses(c echo.Context) error {
	claims := d.authService.GetClaims(&c)
	if err := d.authService.ValidationToken(claims, c.Request().Context()); err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	offlineClasses, err := d.offlineClassService.FindOfflineClasses(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	result := dto.OfflineClassResponses{
		OfflineClasses: *offlineClasses,
		Count:          uint(len(*offlineClasses)),
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get offline classes",
		"data":    result,
	})
}

// OfflineClassBookingPayment implements OfflineClassController
func (d *offlineclassControllerImpl) OfflineClassBookingPayment(c echo.Context) error {
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
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	form, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	src, err := form.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer src.Close()
	body := model.PaymentRequest{
		ID:       uint(id),
		UserID:   uint(claims["user_id"].(float64)),
		FileName: form.Filename,
		File:     src,
	}
	if err := c.Validate(body); err != nil {
		return err
	}
	err = d.offlineClassService.OfflineClassPayment(&body, c.Request().Context())
	if err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "payment success",
	})
}

// SetStatusOfflineClassBooking implements OfflineClassController
func (d *offlineclassControllerImpl) SetStatusOfflineClassBooking(c echo.Context) error {
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
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	var status dto.SetStatusOfflineClassBooking
	if err := c.Bind(&status); err != nil {
		return err
	}
	if err := c.Validate(status); err != nil {
		return err
	}
	status.ID = uint(id)
	if err := d.offlineClassService.SetStatusOfflineClassBooking(&status, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success set status",
	})
}

// TakeOfflineClassBooking implements OfflineClassController
func (d *offlineclassControllerImpl) TakeOfflineClassBooking(c echo.Context) error {
	emailParam := c.QueryParam("email")
	codeParam := c.QueryParam("code")
	code, err := uuid.Parse(codeParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	condition := dto.TakeOfflineClassBooking{
		Email: emailParam,
		Code:  code,
	}
	if err := c.Validate(condition); err != nil {
		return err
	}
	err = d.offlineClassService.TakeOfflineClassBooking(&condition, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success take offline class booking",
	})
}

// UpdateOfflineClass implements OfflineClassController
func (d *offlineclassControllerImpl) UpdateOfflineClass(c echo.Context) error {
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
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	var offlineClass dto.OfflineClassUpdateRequest
	if err := c.Bind(&offlineClass); err != nil {
		return err
	}
	if err := c.Validate(offlineClass); err != nil {
		return err
	}
	offlineClass.ID = uint(id)
	if err := d.offlineClassService.UpdateOfflineClass(&offlineClass, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update offline class",
	})
}

// UpdateOfflineClassBooking implements OfflineClassController
func (d *offlineclassControllerImpl) UpdateOfflineClassBooking(c echo.Context) error {
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
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	var offlineClassBooking dto.OfflineClassBookingUpdateRequest
	if err := c.Bind(&offlineClassBooking); err != nil {
		return err
	}
	if err := c.Validate(offlineClassBooking); err != nil {
		return err
	}
	offlineClassBooking.ID = uint(id)
	if err := d.offlineClassService.UpdateOfflineClassBooking(&offlineClassBooking, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update offline class booking",
	})
}

// UpdateOfflineClassCategory implements OfflineClassController
func (d *offlineclassControllerImpl) UpdateOfflineClassCategory(c echo.Context) error {
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
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	var offlineClassCategory dto.OfflineClassCategoryUpdateRequest
	if err := c.Bind(&offlineClassCategory); err != nil {
		return err
	}
	if err := c.Validate(offlineClassCategory); err != nil {
		return err
	}
	offlineClassCategory.ID = uint(id)
	if err := d.offlineClassService.UpdateOfflineClassCategory(&offlineClassCategory, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update offline class category",
	})
}

func NewOfflineClassController(offlineClassService offlineClassServ.OfflineClassService, authService authServ.AuthService, membersService memberServ.MemberService, notificationServ notificationServ.NotificationService) OfflineClassController {
	return &offlineclassControllerImpl{
		offlineClassService: offlineClassService,
		memberService:       membersService,
		authService:         authService,
		notificationService: notificationServ,
	}
}
