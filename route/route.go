package route

import (
	"github.com/Group10CapstoneProject/Golang/config"
	pkgAuthController "github.com/Group10CapstoneProject/Golang/internal/auth/controller"
	pkgAuthRepostiory "github.com/Group10CapstoneProject/Golang/internal/auth/repository"
	pkgAuthService "github.com/Group10CapstoneProject/Golang/internal/auth/service"
	pkgMemberController "github.com/Group10CapstoneProject/Golang/internal/members/controller"
	pkgMemberRepostiory "github.com/Group10CapstoneProject/Golang/internal/members/repository"
	pkgMemberService "github.com/Group10CapstoneProject/Golang/internal/members/service"
	pkgPaymentMethodController "github.com/Group10CapstoneProject/Golang/internal/paymentMethods/controller"
	pkgPaymentMethodRepostiory "github.com/Group10CapstoneProject/Golang/internal/paymentMethods/repository"
	pkgPaymentMethodService "github.com/Group10CapstoneProject/Golang/internal/paymentMethods/service"
	pkgUserController "github.com/Group10CapstoneProject/Golang/internal/users/controller"
	pkgUserRepostiory "github.com/Group10CapstoneProject/Golang/internal/users/repository"
	pkgUserService "github.com/Group10CapstoneProject/Golang/internal/users/service"
	jwtService "github.com/Group10CapstoneProject/Golang/utils/jwt"
	"github.com/Group10CapstoneProject/Golang/utils/password"
	customValidator "github.com/Group10CapstoneProject/Golang/utils/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	e.Use(middleware.Recover())

	customValidator.NewCustomValidator(e)

	jwtService := jwtService.NewJWTService(config.Env.JWT_SECRET_ACCESS, config.Env.JWT_SECRET_REFRESH)
	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{
				"*",
			},
		},
	))

	api := e.Group("/api")

	// version
	v1 := api.Group("/v1")

	// init user user and auth service
	userRepository := pkgUserRepostiory.NewUserRepository(db)
	authRepository := pkgAuthRepostiory.NewAuthRepository(db)
	authService := pkgAuthService.NewAuthService(authRepository, userRepository, password.Password{}, jwtService)
	userService := pkgUserService.NewUserService(userRepository, password.Password{})
	// create default user (superadmin)
	if err := userService.CreateSuperadmin(); err != nil {
		panic(err)
	}

	// init user and auth controller
	userController := pkgUserController.NewUserController(userService, authService)
	userController.InitRoute(v1)
	authController := pkgAuthController.NewAuthController(authService)
	authController.InitRoute(v1)

	// init member route
	memberRepository := pkgMemberRepostiory.NewMemberRepository(db)
	memberService := pkgMemberService.NewMemberService(memberRepository)
	memberController := pkgMemberController.NewMemberController(memberService, authService)
	memberController.InitRoute(v1)

	// init payment method route
	paymentMethodRepository := pkgPaymentMethodRepostiory.NewPaymentMethodRepository(db)
	paymentMethodService := pkgPaymentMethodService.NewPaymentMethodService(paymentMethodRepository)
	paymentMethodController := pkgPaymentMethodController.NewPaymentMethodController(paymentMethodService, authService)
	paymentMethodController.InitRoute(v1)
}
