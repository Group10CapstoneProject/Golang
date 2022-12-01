package controller

import (
	"net/http"
	"strconv"

	"github.com/Group10CapstoneProject/Golang/config"
	"github.com/Group10CapstoneProject/Golang/constans"
	authServ "github.com/Group10CapstoneProject/Golang/internal/auth/service"
	"github.com/Group10CapstoneProject/Golang/internal/members/dto"
	memberServ "github.com/Group10CapstoneProject/Golang/internal/members/service"
	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type memberControllerImpl struct {
	memberService memberServ.MemberService
	authService   authServ.AuthService
}

// InitRoute implements MemberController
func (d *memberControllerImpl) InitRoute(api *echo.Group) {
	members := api.Group("/members", middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))

	members.POST("", d.CreateMember)
	members.GET("", d.GetMembers)
	members.GET("/user", d.GetMemberUser)
	members.GET("/:id", d.GetMemberDetail)
	members.PUT("/:id", d.UpdateMember)
	members.DELETE("/:id", d.DeleteMember)
	members.POST("/setStatus/:id", d.SetStatusMember)

	memberType := members.Group("/types")
	memberType.POST("", d.CreateMemberType)
	memberType.GET("", d.GetMemberTypes)
	memberType.GET("/:id", d.GetMemberTypeDetail)
	memberType.PUT("/:id", d.UpdateMemberType)
	memberType.DELETE("/:id", d.DeleteMemberType)
}

// CreateMember implements MemberController
func (d *memberControllerImpl) CreateMember(c echo.Context) error {
	claims := d.authService.GetClaims(&c)
	if err := d.authService.ValidationToken(claims, c.Request().Context()); err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	var member dto.MemberStoreRequest
	if err := c.Bind(&member); err != nil {
		return err
	}
	if err := c.Validate(member); err != nil {
		return err
	}
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
	claims := d.authService.GetClaims(&c)
	if err := d.authService.ValidationRole(claims, constans.Role_admin, c.Request().Context()); err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
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
	member, err := d.memberService.FindMemberById(uint(id), c.Request().Context())
	if err != nil {
		if err == myerrors.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get member",
		"data":    member,
	})
}

// GetMemberTypeDetail implements MemberController
func (d *memberControllerImpl) GetMemberTypeDetail(c echo.Context) error {
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
	claims := d.authService.GetClaims(&c)
	if err := d.authService.ValidationToken(claims, c.Request().Context()); err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
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
	claims := d.authService.GetClaims(&c)
	if err := d.authService.ValidationRole(claims, constans.Role_user, c.Request().Context()); err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	members, err := d.memberService.FindMemberByUser(uint(claims["user_id"].(float64)), c.Request().Context())
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
	claims := d.authService.GetClaims(&c)
	if err := d.authService.ValidationRole(claims, constans.Role_admin, c.Request().Context()); err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
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

func NewMemberController(memberService memberServ.MemberService, authService authServ.AuthService) MemberController {
	return &memberControllerImpl{
		memberService: memberService,
		authService:   authService,
	}
}
