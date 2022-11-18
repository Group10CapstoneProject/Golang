package route

import (
	"github.com/Group10CapstoneProject/Golang/config"
	pkgUserController "github.com/Group10CapstoneProject/Golang/internal/users/controller"
	pkgUserRepostiory "github.com/Group10CapstoneProject/Golang/internal/users/repository"
	pkgUserService "github.com/Group10CapstoneProject/Golang/internal/users/service"
	"github.com/Group10CapstoneProject/Golang/utils/password"
	customValidator "github.com/Group10CapstoneProject/Golang/utils/validator"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	e.Use(middleware.Recover())

	e.Validator = &customValidator.CustomValidator{
		Validator: validator.New(),
	}

	api := e.Group("/" + config.Env.API_ENV)

	//version
	v1 := api.Group("/v1")

	// init user controller
	userRepository := pkgUserRepostiory.NewUserRepository(db)
	userService := pkgUserService.NewUserService(userRepository, password.Password{})
	// create default user (superadmin)
	if err := userService.CreateSuperadmin(); err != nil {
		panic(err)
	}
	userController := pkgUserController.NewUserController(userService)
	userController.InitRoute(v1)
}
