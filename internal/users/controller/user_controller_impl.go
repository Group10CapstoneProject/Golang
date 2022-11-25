package controller

import (
	"net/http"

	"github.com/Group10CapstoneProject/Golang/config"
	"github.com/Group10CapstoneProject/Golang/constans"
	authService "github.com/Group10CapstoneProject/Golang/internal/auth/service"
	"github.com/Group10CapstoneProject/Golang/internal/users/dto"
	userService "github.com/Group10CapstoneProject/Golang/internal/users/service"
	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type userControllerImpl struct {
	userService userService.UserService
	authService authService.AuthService
}

// InitRoute implements UserController
func (d *userControllerImpl) InitRoute(api *echo.Group) {
	users := api.Group("/users")

	users.POST("/signup", d.Signup)
	users.GET("", d.GetUsers, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	users.GET("/profile", d.GetUser, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	users.POST("/admin", d.NewAadmin, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	users.GET("/admin", d.GetAdmins, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
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
	claims := d.authService.GetClaims(&c)
	if err := d.authService.ValidationRole(claims, constans.Role_superadmin, c.Request().Context()); err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
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
	claims := d.authService.GetClaims(&c)
	if err := d.authService.ValidationToken(claims, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
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
	claims := d.authService.GetClaims(&c)
	if err := d.authService.ValidationRole(claims, constans.Role_admin, c.Request().Context()); err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
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
	claims := d.authService.GetClaims(&c)
	if err := d.authService.ValidationRole(claims, constans.Role_admin, c.Request().Context()); err != nil {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
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

func NewUserController(userService userService.UserService, authService authService.AuthService) UserController {
	return &userControllerImpl{
		userService: userService,
		authService: authService,
	}
}
