package controller

import (
	"net/http"

	"github.com/Group10CapstoneProject/Golang/internal/users/dto"
	userServ "github.com/Group10CapstoneProject/Golang/internal/users/service"
	"github.com/labstack/echo/v4"
)

type userControllerImpl struct {
	userService userServ.UserService
}

// InitRoute implements UserController
func (d *userControllerImpl) InitRoute(api *echo.Group) {
	api.POST("/signup", d.Signup)
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
func (*userControllerImpl) NewAadmin(c echo.Context) error {
	panic("unimplemented")
}

func NewUserController(userService userServ.UserService) UserController {
	return &userControllerImpl{
		userService: userService,
	}
}
