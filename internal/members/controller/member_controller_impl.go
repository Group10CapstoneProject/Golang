package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Group10CapstoneProject/Golang/constans"
	"github.com/Group10CapstoneProject/Golang/internal/members/dto"
	memberServ "github.com/Group10CapstoneProject/Golang/internal/members/service"
	dtoNotif "github.com/Group10CapstoneProject/Golang/internal/notifications/dto"
	notifServ "github.com/Group10CapstoneProject/Golang/internal/notifications/service"
	"github.com/Group10CapstoneProject/Golang/model"
	jwtServ "github.com/Group10CapstoneProject/Golang/utils/jwt"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"github.com/labstack/echo/v4"
)

type memberControllerImpl struct {
	memberService       memberServ.MemberService
	jwtService          jwtServ.JWTService
	notificationService notifServ.NotificationService
}

// CreateMember implements MemberController
func (d *memberControllerImpl) CreateMember(c echo.Context) error {
	var member dto.MemberStoreRequest
	if err := c.Bind(&member); err != nil {
		return err
	}
	if err := c.Validate(member); err != nil {
		return err
	}
	claims := d.jwtService.GetClaims(&c)
	member.UserID = uint(claims["user_id"].(float64))
	if err := d.memberService.CreateMember(&member, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "new member success created",
	})
}

// CreateMemberType implements MemberController
func (d *memberControllerImpl) CreateMemberType(c echo.Context) error {
	var memberType dto.MemberTypeStoreRequest
	if err := c.Bind(&memberType); err != nil {
		return err
	}
	if err := c.Validate(memberType); err != nil {
		return err
	}
	if err := d.memberService.CreateMemberType(&memberType, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "new member type success created",
	})
}

// DeleteMember implements MemberController
func (d *memberControllerImpl) DeleteMember(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	if err := d.memberService.DeleteMember(uint(id), c.Request().Context()); err != nil {
		if err == myerrors.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete member",
	})
}

// DeleteMemberType implements MemberController
func (d *memberControllerImpl) DeleteMemberType(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	if err := d.memberService.DeleteMemberType(uint(id), c.Request().Context()); err != nil {
		if err == myerrors.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete member type",
	})
}

// GetMemberDetail implements MemberController
func (d *memberControllerImpl) GetMemberDetail(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	claims := d.jwtService.GetClaims(&c)
	member, err := d.memberService.FindMemberById(uint(id), c.Request().Context())
	if err != nil {
		if err == myerrors.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if claims["role"].(string) == constans.Role_superadmin || claims["role"].(string) == constans.Role_admin {
		notif := dtoNotif.NotificationReadRequest{
			TransactionID: uint(id),
			Title:         "Member",
		}
		if err := d.notificationService.ReadNotification(&notif, c.Request().Context()); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get member",
		"data":    member,
	})
}

// GetMemberTypeDetail implements MemberController
func (d *memberControllerImpl) GetMemberTypeDetail(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	memberType, err := d.memberService.FindMemberTypeById(uint(id), c.Request().Context())
	if err != nil {
		if err == myerrors.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get member type",
		"data":    memberType,
	})
}

// GetMemberTypes implements MemberController
func (d *memberControllerImpl) GetMemberTypes(c echo.Context) error {
	memberTypes, err := d.memberService.FindMemberTypes(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get member types",
		"data":    memberTypes,
	})
}

// GetMemberUser implements MemberController
func (d *memberControllerImpl) GetMemberUser(c echo.Context) error {
	claims := d.jwtService.GetClaims(&c)
	userId := claims["user_id"].(float64)
	members, err := d.memberService.FindMemberByUser(uint(userId), c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get members",
		"data":    members,
	})
}

// GetMembers implements MemberController
func (d *memberControllerImpl) GetMembers(c echo.Context) error {
	var query model.Pagination
	query.NewPageQuery(c)

	members, err := d.memberService.FindMembers(&query, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get members",
		"data":    members,
	})
}

// UpdateMember implements MemberController
func (d *memberControllerImpl) UpdateMember(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	var member dto.MemberUpdateRequest
	if err := c.Bind(&member); err != nil {
		return err
	}
	if err := c.Validate(member); err != nil {
		return err
	}
	member.ID = uint(id)
	if err := d.memberService.UpdateMember(&member, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update member",
	})
}

// UpdateMemberType implements MemberController
func (d *memberControllerImpl) UpdateMemberType(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	var memberType dto.MemberTypeUpdateRequest
	if err := c.Bind(&memberType); err != nil {
		return err
	}
	if err := c.Validate(memberType); err != nil {
		return err
	}
	memberType.ID = uint(id)
	if err := d.memberService.UpdateMemberType(&memberType, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update member type",
	})
}

// SetStatusMember implements MemberController
func (d *memberControllerImpl) SetStatusMember(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	var status dto.SetStatusMember
	if err := c.Bind(&status); err != nil {
		return err
	}
	if err := c.Validate(status); err != nil {
		return err
	}
	status.ID = uint(id)
	if err := d.memberService.SetStatusMember(&status, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success set status",
	})
}

// MemberPayment implements MemberController
func (d *memberControllerImpl) MemberPayment(c echo.Context) error {
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
	claims := d.jwtService.GetClaims(&c)
	body := model.PaymentRequest{
		ID:       uint(id),
		UserID:   uint(claims["user_id"].(float64)),
		FileName: form.Filename,
		File:     src,
	}
	if err := c.Validate(body); err != nil {
		return err
	}
	fmt.Println(body.FileName)
	err = d.memberService.MemberPayment(&body, c.Request().Context())
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

func NewMemberController(memberService memberServ.MemberService, notificationService notifServ.NotificationService, jwtService jwtServ.JWTService) MemberController {
	return &memberControllerImpl{
		memberService:       memberService,
		notificationService: notificationService,
		jwtService:          jwtService,
	}
}
