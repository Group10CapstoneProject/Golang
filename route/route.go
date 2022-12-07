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
	pkgOfflineClassController "github.com/Group10CapstoneProject/Golang/internal/offline_classes/controller"
	pkgOfflineClassRepostiory "github.com/Group10CapstoneProject/Golang/internal/offline_classes/repository"
	pkgOfflineClassService "github.com/Group10CapstoneProject/Golang/internal/offline_classes/service"
	pkgOnlineClassController "github.com/Group10CapstoneProject/Golang/internal/online_classes/controller"
	pkgOnlineClassRepostiory "github.com/Group10CapstoneProject/Golang/internal/online_classes/repository"
	pkgOnlineClassService "github.com/Group10CapstoneProject/Golang/internal/online_classes/service"
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
	onlineClassRepository := pkgOnlineClassRepostiory.NewOnlineClassRepository(db)
	offlineClassRepository := pkgOfflineClassRepostiory.NewOfflineClassRepository(db)

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
	onlineClassService := pkgOnlineClassService.NewOnlineClassService(onlineClassRepository, notificationRepository, imagekitService)
	offlineClassService := pkgOfflineClassService.NewOfflineClassService(offlineClassRepository, notificationRepository, imagekitService)

	// init controller
	userController := pkgUserController.NewUserController(userService, authService)
	authController := pkgAuthController.NewAuthController(authService)
	paymentMethodController := pkgPaymentMethodController.NewPaymentMethodController(paymentMethodService, authService)
	memberController := pkgMemberController.NewMemberController(memberService, authService, noticationService)
	fileController := pkgFileController.NewFileController(fileService, authService)
	noticationController := pkgNotificationController.NewNotificationController(noticationService, authService)
	onlineClassController := pkgOnlineClassController.NewOnlineClassController(memberService, authService, noticationService, onlineClassService)
	offlineClassController := pkgOfflineClassController.NewOfflineClassController(offlineClassService, authService, memberService, noticationService)

	// int route
	// auth
	auth := v1.Group("/auth")
	auth.POST("/login", authController.Login)
	auth.POST("/refresh", authController.RefreshToken)
	auth.POST("/logout", authController.Logout, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	admin := auth.Group("/admin")
	admin.POST("/admin/login", authController.LoginAdmin)
	admin.POST("/admin/refresh", authController.RefreshAdminToken)

	// users
	users := v1.Group("/users")
	users.GET("", userController.GetUsers, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	users.POST("/signup", userController.Signup)
	users.GET("/profile", userController.GetUser, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	userAdmin := users.Group("/admin")
	userAdmin.POST("", userController.NewAadmin, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	userAdmin.GET("", userController.GetAdmins, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))

	// payment methods
	paymentMethods := v1.Group("/payment-methods")
	paymentMethods.POST("", paymentMethodController.CreatePaymentMethod, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	paymentMethods.GET("", paymentMethodController.GetPaymentMethods, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	paymentMethodDetails := paymentMethods.Group("/details")
	paymentMethodDetails.GET("/:id", paymentMethodController.GetPaymentMethodDetail, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	paymentMethodDetails.PUT("/:id", paymentMethodController.UpdatePaymentMethod, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	paymentMethodDetails.DELETE("/:id", paymentMethodController.DeletePaymentMethod, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))

	// members
	members := v1.Group("/members")
	members.POST("", memberController.CreateMember, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	members.GET("", memberController.GetMembers, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	members.POST("/set-status/:id", memberController.SetStatusMember, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	members.POST("/pay/:id", memberController.MemberPayment, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	members.GET("/user", memberController.GetMemberUser, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	memberDetails := members.Group("/details")
	memberDetails.GET("/:id", memberController.GetMemberDetail, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	memberDetails.PUT("/:id", memberController.UpdateMember, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	memberDetails.DELETE("/:id", memberController.DeleteMember, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	// member types
	memberTypes := members.Group("/types")
	memberTypes.POST("", memberController.CreateMemberType, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	memberTypes.GET("", memberController.GetMemberTypes, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	memberTypeDetails := memberTypes.Group("/details")
	memberTypeDetails.GET("/:id", memberController.GetMemberTypeDetail, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	memberTypeDetails.PUT("/:id", memberController.UpdateMemberType, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	memberTypeDetails.DELETE("/:id", memberController.DeleteMemberType, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))

	// files
	files := v1.Group("/files")
	files.POST("/upload", fileController.Upload, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))

	// notifications
	notifications := v1.Group("/notifications")
	notifications.GET("", noticationController.GetNotifications, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))

	// online classes
	onlineClasses := v1.Group("/online-classes")
	onlineClasses.POST("", onlineClassController.CreateOnlineClass, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	onlineClasses.GET("", onlineClassController.GetOnlineClasses, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	onlineClassDetails := onlineClasses.Group("/details")
	onlineClassDetails.GET("/:id", onlineClassController.GetOnlineClassDetail, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	onlineClassDetails.PUT("/:id", onlineClassController.UpdateOnlineClass, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	onlineClassDetails.DELETE("/:id", onlineClassController.DeleteOnlineClass, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	// online class booking
	onlineClassBooking := onlineClasses.Group("/bookings")
	onlineClassBooking.POST("", onlineClassController.CreateOnlineClassBooking, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	onlineClassBooking.GET("", onlineClassController.GetOnlineClassBookings, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	onlineClassBooking.GET("/user", onlineClassController.GetOnlineClassBookingUser, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	onlineClassBooking.POST("/set-status/:id", onlineClassController.SetStatusOnlineClassBooking, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	onlineClassBooking.POST("/pay/:id", onlineClassController.OnlineClassBookingPayment, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	onlineClassBookingDetails := onlineClassBooking.Group("/details")
	onlineClassBookingDetails.GET("/:id", onlineClassController.GetOnlineClassBookingDetail, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	onlineClassBookingDetails.PUT("/:id", onlineClassController.UpdateOnlineClassBooking, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	onlineClassBookingDetails.DELETE("/:id", onlineClassController.DeleteOnlineClassBooking, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	// online class category
	onlineClassCategory := onlineClasses.Group("/categories")
	onlineClassCategory.POST("", onlineClassController.CreateOnlineClassCategory, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	onlineClassCategory.GET("", onlineClassController.GetOnlineClassCategories, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	onlineClassCategoryDetails := onlineClassCategory.Group("/details")
	onlineClassCategoryDetails.GET("/:id", onlineClassController.GetOnlineClassCategoryDetail, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	onlineClassCategoryDetails.PUT("/:id", onlineClassController.UpdateOnlineClassCategory, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	onlineClassCategoryDetails.DELETE("/:id", onlineClassController.DeleteOnlineClassCategory, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))

	// offline classes
	offlineClasses := v1.Group("/offline-classes")
	offlineClasses.POST("", offlineClassController.CreateOfflineClass, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	offlineClasses.GET("", offlineClassController.GetOfflineClasses, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	offlineClassDetails := offlineClasses.Group("/details")
	offlineClassDetails.GET("/:id", offlineClassController.GetOfflineClassDetail, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	offlineClassDetails.PUT("/:id", offlineClassController.UpdateOfflineClass, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	offlineClassDetails.DELETE("/:id", offlineClassController.DeleteOfflineClass, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	// offline class booking
	offlineClassBooking := offlineClasses.Group("/bookings")
	offlineClassBooking.POST("", offlineClassController.CreateOfflineClassBooking, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	offlineClassBooking.GET("", offlineClassController.GetOfflineClassBookings, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	offlineClassBooking.GET("/user", offlineClassController.GetOfflineClassBookingUser, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	offlineClassBooking.POST("/set-status/:id", offlineClassController.SetStatusOfflineClassBooking, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	offlineClassBooking.POST("/pay/:id", offlineClassController.OfflineClassBookingPayment, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	offlineClassBookingTake := offlineClassBooking.Group("/take")
	offlineClassBookingTake.GET("", offlineClassController.CheckOfflineClassBooking, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	offlineClassBookingTake.POST("", offlineClassController.TakeOfflineClassBooking, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	offlineClassBookingDetails := offlineClassBooking.Group("/details")
	offlineClassBookingDetails.GET("/:id", offlineClassController.GetOfflineClassBookingDetail, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	offlineClassBookingDetails.PUT("/:id", offlineClassController.UpdateOfflineClassBooking, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	offlineClassBookingDetails.DELETE("/:id", offlineClassController.DeleteOfflineClassBooking, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	// offline class category
	offlineClassCategory := offlineClasses.Group("/categories")
	offlineClassCategory.POST("", offlineClassController.CreateOfflineClassCategory, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	offlineClassCategory.GET("", offlineClassController.GetOfflineClassCategories, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	offlineClassCategoryDetails := offlineClassCategory.Group("/details")
	offlineClassCategoryDetails.GET("/:id", offlineClassController.GetOfflineClassCategoryDetail, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	offlineClassCategoryDetails.PUT("/:id", offlineClassController.UpdateOfflineClassCategory, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))
	offlineClassCategoryDetails.DELETE("/:id", offlineClassController.DeleteOfflineClassCategory, middleware.JWT([]byte(config.Env.JWT_SECRET_ACCESS)))

	// create default user (superadmin)
	if err := userService.CreateSuperadmin(); err != nil {
		panic(err)
	}
}
