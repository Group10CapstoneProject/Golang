package route

import (
	"github.com/Group10CapstoneProject/Golang/config"
	pkgAuthController "github.com/Group10CapstoneProject/Golang/internal/auth/controller"
	pkgAuthRepostiory "github.com/Group10CapstoneProject/Golang/internal/auth/repository"
	pkgAuthService "github.com/Group10CapstoneProject/Golang/internal/auth/service"
	pkgFileController "github.com/Group10CapstoneProject/Golang/internal/file/controller"
	pkgFileService "github.com/Group10CapstoneProject/Golang/internal/file/service"
	pkgMemberController "github.com/Group10CapstoneProject/Golang/internal/members/controller"
	pkgMemberRepostiory "github.com/Group10CapstoneProject/Golang/internal/members/repository"
	pkgMemberService "github.com/Group10CapstoneProject/Golang/internal/members/service"
	pkgNotificationController "github.com/Group10CapstoneProject/Golang/internal/notifications/controller"
	pkgNotificationRepostiory "github.com/Group10CapstoneProject/Golang/internal/notifications/repository"
	pkgNotificationService "github.com/Group10CapstoneProject/Golang/internal/notifications/service"
	pkgPaymentMethodController "github.com/Group10CapstoneProject/Golang/internal/paymentMethods/controller"
	pkgPaymentMethodRepostiory "github.com/Group10CapstoneProject/Golang/internal/paymentMethods/repository"
	pkgPaymentMethodService "github.com/Group10CapstoneProject/Golang/internal/paymentMethods/service"
	pkgUserController "github.com/Group10CapstoneProject/Golang/internal/users/controller"
	pkgUserRepostiory "github.com/Group10CapstoneProject/Golang/internal/users/repository"
	pkgUserService "github.com/Group10CapstoneProject/Golang/internal/users/service"
	"github.com/Group10CapstoneProject/Golang/utils/imgkit"
	jwtService "github.com/Group10CapstoneProject/Golang/utils/jwt"
	"github.com/Group10CapstoneProject/Golang/utils/password"
	customValidator "github.com/Group10CapstoneProject/Golang/utils/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{
				"*",
			},
		},
	))
	customValidator.NewCustomValidator(e)

	api := e.Group("/api")

	// version
	v1 := api.Group("/v1")

	// init internal repository
	userRepository := pkgUserRepostiory.NewUserRepository(db)
	authRepository := pkgAuthRepostiory.NewAuthRepository(db)
	notificationRepository := pkgNotificationRepostiory.NewNotificationRepository(db)
	paymentMethodRepository := pkgPaymentMethodRepostiory.NewPaymentMethodRepository(db)
	memberRepository := pkgMemberRepostiory.NewMemberRepository(db)

	// init utils service
	jwtService := jwtService.NewJWTService(config.Env.JWT_SECRET_ACCESS, config.Env.JWT_SECRET_REFRESH)
	imagekitService := imgkit.NewImageKitService(config.Env.IMAGEKIT_PRIVKEY, config.Env.IMAGEKIT_PUBKEY)
	passwordService := password.NewPasswordService()

	// init internal service
	authService := pkgAuthService.NewAuthService(authRepository, userRepository, passwordService, jwtService)
	userService := pkgUserService.NewUserService(userRepository, passwordService)
	paymentMethodService := pkgPaymentMethodService.NewPaymentMethodService(paymentMethodRepository)
	noticationService := pkgNotificationService.NewNotificationService(notificationRepository)
	memberService := pkgMemberService.NewMemberService(memberRepository, imagekitService, notificationRepository)
	fileService := pkgFileService.NewFileService(imagekitService)

	// init controller
	userController := pkgUserController.NewUserController(userService, authService)
	authController := pkgAuthController.NewAuthController(authService)
	paymentMethodController := pkgPaymentMethodController.NewPaymentMethodController(paymentMethodService, authService)
	memberController := pkgMemberController.NewMemberController(memberService, authService, noticationService)
	fileController := pkgFileController.NewFileController(fileService, authService)
	noticationController := pkgNotificationController.NewNotificationController(noticationService, authService)

	// int route
	// auth
	auth := v1.Group("/auth")
	auth.POST("/login", authController.Login)
	auth.POST("/refresh", authController.RefreshToken)
	auth.POST("/admin/login", authController.LoginAdmin)
	auth.POST("/admin/refresh", authController.RefreshAdminToken)
	auth.POST("/logout", authController.Logout, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	// users
	users := v1.Group("/users")
	users.POST("/signup", userController.Signup)
	users.GET("", userController.GetUsers, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	users.GET("/profile", userController.GetUser, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	users.POST("/admin", userController.NewAadmin, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	users.GET("/admin", userController.GetAdmins, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	// payment methods
	paymentMethods := v1.Group("/paymentMethods", middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	paymentMethods.POST("", paymentMethodController.CreatePaymentMethod)
	paymentMethods.GET("", paymentMethodController.GetPaymentMethods)
	paymentMethods.GET("/:id", paymentMethodController.GetPaymentMethodDetail)
	paymentMethods.PUT("/:id", paymentMethodController.UpdatePaymentMethod)
	paymentMethods.DELETE("/:id", paymentMethodController.DeletePaymentMethod)
	// members
	members := v1.Group("/members", middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	members.POST("", memberController.CreateMember)
	members.GET("", memberController.GetMembers)
	members.GET("/user", memberController.GetMemberUser)
	members.GET("/:id", memberController.GetMemberDetail)
	members.PUT("/:id", memberController.UpdateMember)
	members.DELETE("/:id", memberController.DeleteMember)
	members.POST("/setStatus/:id", memberController.SetStatusMember)
	members.POST("/pay/:id", memberController.MemberPayment)
	// member types
	memberTypes := members.Group("/types")
	memberTypes.POST("", memberController.CreateMemberType)
	memberTypes.GET("", memberController.GetMemberTypes)
	memberTypes.GET("/:id", memberController.GetMemberTypeDetail)
	memberTypes.PUT("/:id", memberController.UpdateMemberType)
	memberTypes.DELETE("/:id", memberController.DeleteMemberType)
	// files
	files := v1.Group("/files", middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	files.POST("/upload", fileController.Upload)
	// notifications
	notifications := v1.Group("/notifications", middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	notifications.GET("", noticationController.GetNotifications)

	// create default user (superadmin)
	if err := userService.CreateSuperadmin(); err != nil {
		panic(err)
	}
}
