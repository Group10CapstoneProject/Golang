package route

import (
	"github.com/Group10CapstoneProject/Golang/config"
	pkgAuthController "github.com/Group10CapstoneProject/Golang/internal/auth/controller"
	pkgAuthRepostiory "github.com/Group10CapstoneProject/Golang/internal/auth/repository"
	pkgAuthService "github.com/Group10CapstoneProject/Golang/internal/auth/service"
	pkgUserController "github.com/Group10CapstoneProject/Golang/internal/users/controller"
	pkgUserRepostiory "github.com/Group10CapstoneProject/Golang/internal/users/repository"
	pkgUserService "github.com/Group10CapstoneProject/Golang/internal/users/service"
	jwtService "github.com/Group10CapstoneProject/Golang/utils/jwt"
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
	jwtService := jwtService.NewJWTService(config.Env.JWT_SECRET_ACCESS, config.Env.JWT_SECRET_REFRESH)

	api := e.Group("/api")
	api.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{
				"*",
			},
			AllowHeaders: []string{
				"*",
			},
		},
	))

	//version
	v1 := api.Group("/v1")
	protect := v1.Group("")
	protect.Use(middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))

	// init user controller
	userRepository := pkgUserRepostiory.NewUserRepository(db)
	userService := pkgUserService.NewUserService(userRepository, password.Password{})
	// create default user (superadmin)
	if err := userService.CreateSuperadmin(); err != nil {
		panic(err)
	}
	userController := pkgUserController.NewUserController(userService)
	userController.InitRoute(v1)

	// init auth controller
	authRepository := pkgAuthRepostiory.NewAuthRepository(db)
	authService := pkgAuthService.NewAuthService(authRepository, userRepository, password.Password{}, jwtService)
	authController := pkgAuthController.NewAuthController(authService)
	authController.InitRoute(v1, protect)
}
