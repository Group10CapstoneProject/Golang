package controller

import (
	"net/http"
	"strconv"

	"github.com/Group10CapstoneProject/Golang/constans"
	"github.com/Group10CapstoneProject/Golang/internal/users/dto"
	userService "github.com/Group10CapstoneProject/Golang/internal/users/service"
	"github.com/Group10CapstoneProject/Golang/model"
	jwtServ "github.com/Group10CapstoneProject/Golang/utils/jwt"
	"github.com/labstack/echo/v4"
)

type userControllerImpl struct {
	userService userService.UserService
	jwtService  jwtServ.JWTService
}

// DeleteAdmin implements UserController
func (d *userControllerImpl) DeleteAdmin(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not fount")
	}
	idUint := uint(id)
	admin, err := d.userService.FindUser(&idUint, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if admin.Role != constans.Role_admin {
		return echo.NewHTTPError(http.StatusBadRequest, "record not found")
	}
	err = d.userService.DeleteUser(&idUint, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete admin",
	})
}

// GetAdminDetail implements UserController
func (d *userControllerImpl) GetAdminDetail(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not fount")
	}
	idUint := uint(id)
	admin, err := d.userService.FindUser(&idUint, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if admin.Role != constans.Role_admin {
		return echo.NewHTTPError(http.StatusBadRequest, "record not found")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get admin",
		"data":    admin,
	})
}

// UpdateAdmin implements UserController
func (d *userControllerImpl) UpdateAdmin(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not fount")
	}
	idUint := uint(id)
	admin, err := d.userService.FindUser(&idUint, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if admin.Role != constans.Role_admin {
		return echo.NewHTTPError(http.StatusBadRequest, "record not found")
	}
	var user dto.UpdateUser
	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := c.Validate(user); err != nil {
		return err
	}
	user.ID = idUint
	if err := d.userService.UpdateUser(&user, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update admin",
	})
}

// Signup implements UserController
func (d *userControllerImpl) Signup(c echo.Context) error {
	var user dto.NewUser
	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := c.Validate(user); err != nil {
		return err
	}
	if err := d.userService.CreateUser(&user, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "sign up success",
	})
}

// NewAadmin implements UserController
func (d *userControllerImpl) NewAadmin(c echo.Context) error {
	var user dto.NewUser
	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := c.Validate(user); err != nil {
		return err
	}
	if err := d.userService.CreateAdmin(&user, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "new admin success created",
	})
}

// GetUser implements UserController
func (d *userControllerImpl) GetUser(c echo.Context) error {
	claims := d.jwtService.GetClaims(&c)
	userId := uint(claims["user_id"].(float64))
	user, err := d.userService.FindUser(&userId, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get profile",
		"data":    user,
	})
}

// GetUsers implements UserController
func (d *userControllerImpl) GetUsers(c echo.Context) error {
	var query model.Pagination
	query.NewPageQuery(c)

	users, err := d.userService.FindUsers(query, constans.Role_user, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get users",
		"data":    users,
	})
}

// GetAdmins implements UserController
func (d *userControllerImpl) GetAdmins(c echo.Context) error {
	var query model.Pagination
	query.NewPageQuery(c)

	users, err := d.userService.FindUsers(query, constans.Role_admin, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get admins",
		"data":    users,
	})
}

func NewUserController(userService userService.UserService, jwtService jwtServ.JWTService) UserController {
	return &userControllerImpl{
		userService: userService,
		jwtService:  jwtService,
	}
}
