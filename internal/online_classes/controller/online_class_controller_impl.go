package controller

import (
	"net/http"
	"strconv"

	"github.com/Group10CapstoneProject/Golang/constans"
	memberServ "github.com/Group10CapstoneProject/Golang/internal/members/service"
	dtoNotif "github.com/Group10CapstoneProject/Golang/internal/notifications/dto"
	notifServ "github.com/Group10CapstoneProject/Golang/internal/notifications/service"
	"github.com/Group10CapstoneProject/Golang/internal/online_classes/dto"
	onlineClassServ "github.com/Group10CapstoneProject/Golang/internal/online_classes/service"
	"github.com/Group10CapstoneProject/Golang/model"
	jwtServ "github.com/Group10CapstoneProject/Golang/utils/jwt"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"github.com/labstack/echo/v4"
)

type onlineClassControllerImpl struct {
	memberService       memberServ.MemberService
	onlineClassService  onlineClassServ.OnlineClassService
	jwtService          jwtServ.JWTService
	notificationService notifServ.NotificationService
}

// CreateOnlineClass implements OnlineClassController
func (d *onlineClassControllerImpl) CreateOnlineClass(c echo.Context) error {
	var onlineClass dto.OnlineClassStoreRequest
	if err := c.Bind(&onlineClass); err != nil {
		return err
	}
	if err := c.Validate(onlineClass); err != nil {
		return err
	}
	if err := d.onlineClassService.CreateOnlineClass(&onlineClass, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "new online class success created",
	})
}

// CreateOnlineClassBooking implements OnlineClassController
func (d *onlineClassControllerImpl) CreateOnlineClassBooking(c echo.Context) error {
	var onlineClassBooking dto.OnlineClassBookingStoreRequest
	if err := c.Bind(&onlineClassBooking); err != nil {
		return err
	}
	if err := c.Validate(onlineClassBooking); err != nil {
		return err
	}
	claims := d.jwtService.GetClaims(&c)
	onlineClassBooking.UserID = uint(claims["user_id"].(float64))
	if err := d.onlineClassService.CreateOnlineClassBooking(&onlineClassBooking, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "new online class booking success created",
	})
}

// CreateOnlineClassCategory implements OnlineClassController
func (d *onlineClassControllerImpl) CreateOnlineClassCategory(c echo.Context) error {
	var onlineClassCategory dto.OnlineClassCategoryStoreRequest
	if err := c.Bind(&onlineClassCategory); err != nil {
		return err
	}
	if err := c.Validate(onlineClassCategory); err != nil {
		return err
	}
	if err := d.onlineClassService.CreateOnlineClassCategory(&onlineClassCategory, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "new online class category success created",
	})
}

// DeleteOnlineClass implements OnlineClassController
func (d *onlineClassControllerImpl) DeleteOnlineClass(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	if err := d.onlineClassService.DeleteOnlineClass(uint(id), c.Request().Context()); err != nil {
		if err == myerrors.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete online class",
	})
}

// DeleteOnlineClassBooking implements OnlineClassController
func (d *onlineClassControllerImpl) DeleteOnlineClassBooking(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	if err := d.onlineClassService.DeleteOnlineClassBooking(uint(id), c.Request().Context()); err != nil {
		if err == myerrors.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete online class booking",
	})
}

// DeleteOnlineClassCategory implements OnlineClassController
func (d *onlineClassControllerImpl) DeleteOnlineClassCategory(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	if err := d.onlineClassService.DeleteOnlineClassCategory(uint(id), c.Request().Context()); err != nil {
		if err == myerrors.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete online class category",
	})
}

// GetOnlineClassBookingDetail implements OnlineClassController
func (d *onlineClassControllerImpl) GetOnlineClassBookingDetail(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	onlineClassBooking, err := d.onlineClassService.FindOnlineClassBookingById(uint(id), c.Request().Context())
	if err != nil {
		if err == myerrors.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	claims := d.jwtService.GetClaims(&c)
	if claims["role"].(string) == constans.Role_superadmin || claims["role"].(string) == constans.Role_admin {
		notif := dtoNotif.NotificationReadRequest{
			TransactionID: uint(id),
			Title:         "Online Class",
		}
		if err := d.notificationService.ReadNotification(&notif, c.Request().Context()); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get online class booking",
		"data":    onlineClassBooking,
	})
}

// GetOnlineClassBookingUser implements OnlineClassController
func (d *onlineClassControllerImpl) GetOnlineClassBookingUser(c echo.Context) error {
	claims := d.jwtService.GetClaims(&c)
	onlineClassBooking, err := d.onlineClassService.FindOnlineClassBookingByUser(uint(claims["user_id"].(float64)), c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get online class booking",
		"data":    onlineClassBooking,
	})
}

// GetOnlineClassBookings implements OnlineClassController
func (d *onlineClassControllerImpl) GetOnlineClassBookings(c echo.Context) error {
	var query model.Pagination
	query.NewPageQuery(c)

	onlineClassBookings, err := d.onlineClassService.FindOnlineClassBookings(&query, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get online class bookings",
		"data":    onlineClassBookings,
	})
}

// GetOnlineClassCategories implements OnlineClassController
func (d *onlineClassControllerImpl) GetOnlineClassCategories(c echo.Context) error {
	onlineClassCategories, err := d.onlineClassService.FindOnlineClassCategories(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get online class categories",
		"data":    onlineClassCategories,
	})
}

// GetOnlineClassCategoryDetail implements OnlineClassController
func (d *onlineClassControllerImpl) GetOnlineClassCategoryDetail(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	onlineClassCategory, err := d.onlineClassService.FindOnlineClassCategoryById(uint(id), c.Request().Context())
	if err != nil {
		if err == myerrors.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get online calss category",
		"data":    onlineClassCategory,
	})
}

// GetOnlineClassDetail implements OnlineClassController
func (d *onlineClassControllerImpl) GetOnlineClassDetail(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	onlineClass, err := d.onlineClassService.FindOnlineClassById(uint(id), c.Request().Context())
	if err != nil {
		if err == myerrors.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	onlineClass.AccessClass = true
	claims := d.jwtService.GetClaims(&c)

	if claims["role"].(string) == constans.Role_user {
		memberUser, err := d.memberService.FindMemberByUser(uint(claims["user_id"].(float64)), c.Request().Context())
		if err != nil {
			if err != myerrors.ErrRecordNotFound {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			onlineClass.AccessClass = false
		} else {
			onlineClass.AccessClass = memberUser.MemberType.AccessOnlineClass
		}
		if !onlineClass.AccessClass {
			onlineClass.AccessClass, err = d.onlineClassService.CheckAccessOnlineClass(uint(claims["user_id"].(float64)), uint(id), c.Request().Context())
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get online class",
		"data":    onlineClass,
	})
}

// GetOnlineClasses implements OnlineClassController
func (d *onlineClassControllerImpl) GetOnlineClasses(c echo.Context) error {
	onlineClasses, err := d.onlineClassService.FindOnlineClasses(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get online classes",
		"data":    onlineClasses,
	})
}

// OnlineClassBookingPayment implements OnlineClassController
func (d *onlineClassControllerImpl) OnlineClassBookingPayment(c echo.Context) error {
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
	claims := d.jwtService.GetClaims(&c)
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
	err = d.onlineClassService.OnlineClassPayment(&body, c.Request().Context())
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

// SetStatusOnlineClassBooking implements OnlineClassController
func (d *onlineClassControllerImpl) SetStatusOnlineClassBooking(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	var status dto.SetStatusOnlineClassBooking
	if err := c.Bind(&status); err != nil {
		return err
	}
	if err := c.Validate(status); err != nil {
		return err
	}
	status.ID = uint(id)
	if err := d.onlineClassService.SetStatusOnlineClassBooking(&status, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success set status",
	})
}

// UpdateOnlineClass implements OnlineClassController
func (d *onlineClassControllerImpl) UpdateOnlineClass(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	var onlineClass dto.OnlineClassUpdateRequest
	if err := c.Bind(&onlineClass); err != nil {
		return err
	}
	if err := c.Validate(onlineClass); err != nil {
		return err
	}
	onlineClass.ID = uint(id)
	if err := d.onlineClassService.UpdateOnlineClass(&onlineClass, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update online class",
	})
}

// UpdateOnlineClassBooking implements OnlineClassController
func (d *onlineClassControllerImpl) UpdateOnlineClassBooking(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	var onlineClassBooking dto.OnlineClassBookingUpdateRequest
	if err := c.Bind(&onlineClassBooking); err != nil {
		return err
	}
	if err := c.Validate(onlineClassBooking); err != nil {
		return err
	}
	onlineClassBooking.ID = uint(id)
	if err := d.onlineClassService.UpdateOnlineClassBooking(&onlineClassBooking, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update online class booking",
	})
}

// UpdateOnlineClassCategory implements OnlineClassController
func (d *onlineClassControllerImpl) UpdateOnlineClassCategory(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	var onlineClassCategory dto.OnlineClassCategoryUpdateRequest
	if err := c.Bind(&onlineClassCategory); err != nil {
		return err
	}
	if err := c.Validate(onlineClassCategory); err != nil {
		return err
	}
	onlineClassCategory.ID = uint(id)
	if err := d.onlineClassService.UpdateOnlineClassCategory(&onlineClassCategory, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update online class category",
	})
}

func NewOnlineClassController(memberService memberServ.MemberService, jwtService jwtServ.JWTService, notificationService notifServ.NotificationService, onlineClassService onlineClassServ.OnlineClassService) OnlineClassController {
	return &onlineClassControllerImpl{
		memberService:       memberService,
		onlineClassService:  onlineClassService,
		jwtService:          jwtService,
		notificationService: notificationService,
	}
}
