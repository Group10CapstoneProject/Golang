package controller

import (
	"net/http"

	authService "github.com/Group10CapstoneProject/Golang/internal/auth/service"
	"github.com/Group10CapstoneProject/Golang/internal/users/dto"
	userService "github.com/Group10CapstoneProject/Golang/internal/users/service"
	"github.com/labstack/echo/v4"
)

type userControllerImpl struct {
	userService userService.UserService
	authService authService.AuthService
}

// InitRoute implements UserController
func (d *userControllerImpl) InitRoute(api *echo.Group, protect *echo.Group) {
	users := api.Group("/users")
	protectUsers := protect.Group("/users")

	users.POST("/signup", d.Signup)
	protectUsers.GET("/profile", d.GetUser)
	protectUsers.GET("", d.GetUser)
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
	panic("unimplemented")
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
func (*userControllerImpl) GetUsers(c echo.Context) error {
	panic("unimplemented")
}

func NewUserController(userService userService.UserService, authService authService.AuthService) UserController {
	return &userControllerImpl{
		userService: userService,
		authService: authService,
	}
}
