package route

import (
	"github.com/Group10CapstoneProject/Golang/config"
	"github.com/Group10CapstoneProject/Golang/constans"
	pkgAuthController "github.com/Group10CapstoneProject/Golang/internal/auth/controller"
	pkgAuthRepostiory "github.com/Group10CapstoneProject/Golang/internal/auth/repository"
	pkgAuthService "github.com/Group10CapstoneProject/Golang/internal/auth/service"
	pkgFileController "github.com/Group10CapstoneProject/Golang/internal/file/controller"
	pkgFileService "github.com/Group10CapstoneProject/Golang/internal/file/service"
	pkgHistoryController "github.com/Group10CapstoneProject/Golang/internal/history/controller"
	pkgHistoryService "github.com/Group10CapstoneProject/Golang/internal/history/service"
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
	pkgPaymentMethodController "github.com/Group10CapstoneProject/Golang/internal/payment_methods/controller"
	pkgPaymentMethodRepostiory "github.com/Group10CapstoneProject/Golang/internal/payment_methods/repository"
	pkgPaymentMethodService "github.com/Group10CapstoneProject/Golang/internal/payment_methods/service"
	pkgUserController "github.com/Group10CapstoneProject/Golang/internal/users/controller"
	pkgUserRepostiory "github.com/Group10CapstoneProject/Golang/internal/users/repository"
	pkgUserService "github.com/Group10CapstoneProject/Golang/internal/users/service"
	customMiddleware "github.com/Group10CapstoneProject/Golang/middleware"
	"github.com/Group10CapstoneProject/Golang/model"
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
	md := customMiddleware.NewCustomMiddleware(db, config.Env.JWT_SECRET_ACCESS)
	allAccess := ""
	roleSuperadmin := constans.Role_superadmin
	roleAdmin := constans.Role_admin
	roleUser := constans.Role_user

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
	historyService := pkgHistoryService.NewHistoryService(memberRepository, onlineClassRepository, offlineClassRepository)

	// init controller
	userController := pkgUserController.NewUserController(userService, jwtService)
	authController := pkgAuthController.NewAuthController(authService)
	paymentMethodController := pkgPaymentMethodController.NewPaymentMethodController(paymentMethodService, jwtService)
	memberController := pkgMemberController.NewMemberController(memberService, noticationService, jwtService)
	fileController := pkgFileController.NewFileController(fileService)
	noticationController := pkgNotificationController.NewNotificationController(noticationService)
	onlineClassController := pkgOnlineClassController.NewOnlineClassController(memberService, jwtService, noticationService, onlineClassService)
	offlineClassController := pkgOfflineClassController.NewOfflineClassController(offlineClassService, jwtService, memberService, noticationService)
	historyController := pkgHistoryController.NewHistoryController(historyService, jwtService)

	// int route
	// auth
	auth := v1.Group("/auth")
	auth.POST("/login", authController.Login)
	auth.POST("/refresh", authController.RefreshToken)
	auth.POST("/logout", authController.Logout, md.CustomJWTWithConfig(allAccess))
	admin := auth.Group("/admin")
	admin.POST("/login", authController.LoginAdmin)
	admin.POST("/refresh", authController.RefreshAdminToken)

	// users
	users := v1.Group("/users")
	users.GET("", userController.GetUsers, md.CustomJWTWithConfig(roleAdmin))
	users.POST("/signup", userController.Signup)
	users.GET("/profile", userController.GetUser, md.CustomJWTWithConfig(allAccess))
	userAdmin := users.Group("/admin")
	userAdmin.POST("", userController.NewAadmin, md.CustomJWTWithConfig(roleSuperadmin))
	userAdmin.GET("", userController.GetAdmins, md.CustomJWTWithConfig(roleSuperadmin))

	// payment methods
	paymentMethods := v1.Group("/payment-methods")
	paymentMethods.POST("", paymentMethodController.CreatePaymentMethod, md.CustomJWTWithConfig(roleAdmin))
	paymentMethods.GET("", paymentMethodController.GetPaymentMethods, md.CustomJWTWithConfig(allAccess))
	paymentMethodDetails := paymentMethods.Group("/details")
	paymentMethodDetails.GET("/:id", paymentMethodController.GetPaymentMethodDetail, md.CustomJWTWithConfig(allAccess))
	paymentMethodDetails.PUT("/:id", paymentMethodController.UpdatePaymentMethod, md.CustomJWTWithConfig(roleAdmin))
	paymentMethodDetails.DELETE("/:id", paymentMethodController.DeletePaymentMethod, md.CustomJWTWithConfig(roleAdmin))

	// members
	members := v1.Group("/members")
	members.POST("", memberController.CreateMember, md.CustomJWTWithConfig(roleUser))
	members.GET("", memberController.GetMembers, md.CustomJWTWithConfig(roleAdmin))
	members.POST("/set-status/:id", memberController.SetStatusMember, md.CustomJWTWithConfig(roleAdmin))
	members.POST("/pay/:id", memberController.MemberPayment, md.CustomJWTWithConfig(roleUser))
	members.GET("/user", memberController.GetMemberUser, md.CustomJWTWithConfig(roleUser))
	memberDetails := members.Group("/details")
	memberDetails.GET("/:id", memberController.GetMemberDetail, md.CustomJWTWithConfig(allAccess))
	memberDetails.PUT("/:id", memberController.UpdateMember, md.CustomJWTWithConfig(roleAdmin))
	memberDetails.DELETE("/:id", memberController.DeleteMember, md.CustomJWTWithConfig(roleAdmin))
	// member types
	memberTypes := members.Group("/types")
	memberTypes.POST("", memberController.CreateMemberType, md.CustomJWTWithConfig(roleAdmin))
	memberTypes.GET("", memberController.GetMemberTypes, md.CustomJWTWithConfig(allAccess))
	memberTypeDetails := memberTypes.Group("/details")
	memberTypeDetails.GET("/:id", memberController.GetMemberTypeDetail, md.CustomJWTWithConfig(allAccess))
	memberTypeDetails.PUT("/:id", memberController.UpdateMemberType, md.CustomJWTWithConfig(roleAdmin))
	memberTypeDetails.DELETE("/:id", memberController.DeleteMemberType, md.CustomJWTWithConfig(roleAdmin))

	// files
	files := v1.Group("/files")
	files.POST("/upload", fileController.Upload, md.CustomJWTWithConfig(roleAdmin))

	// notifications
	notifications := v1.Group("/notifications")
	notifications.GET("", noticationController.GetNotifications, md.CustomJWTWithConfig(roleAdmin))

	// online classes
	onlineClasses := v1.Group("/online-classes")
	onlineClasses.POST("", onlineClassController.CreateOnlineClass, md.CustomJWTWithConfig(roleAdmin))
	onlineClasses.GET("", onlineClassController.GetOnlineClasses, md.CustomJWTWithConfig(allAccess))
	onlineClassDetails := onlineClasses.Group("/details")
	onlineClassDetails.GET("/:id", onlineClassController.GetOnlineClassDetail, md.CustomJWTWithConfig(allAccess))
	onlineClassDetails.PUT("/:id", onlineClassController.UpdateOnlineClass, md.CustomJWTWithConfig(roleAdmin))
	onlineClassDetails.DELETE("/:id", onlineClassController.DeleteOnlineClass, md.CustomJWTWithConfig(roleAdmin))
	// online class booking
	onlineClassBooking := onlineClasses.Group("/bookings")
	onlineClassBooking.POST("", onlineClassController.CreateOnlineClassBooking, md.CustomJWTWithConfig(roleUser))
	onlineClassBooking.GET("", onlineClassController.GetOnlineClassBookings, md.CustomJWTWithConfig(roleAdmin))
	onlineClassBooking.GET("/user", onlineClassController.GetOnlineClassBookingUser, md.CustomJWTWithConfig(roleUser))
	onlineClassBooking.POST("/set-status/:id", onlineClassController.SetStatusOnlineClassBooking, md.CustomJWTWithConfig(roleAdmin))
	onlineClassBooking.POST("/pay/:id", onlineClassController.OnlineClassBookingPayment, md.CustomJWTWithConfig(roleUser))
	onlineClassBookingDetails := onlineClassBooking.Group("/details")
	onlineClassBookingDetails.GET("/:id", onlineClassController.GetOnlineClassBookingDetail, md.CustomJWTWithConfig(roleUser))
	onlineClassBookingDetails.PUT("/:id", onlineClassController.UpdateOnlineClassBooking, md.CustomJWTWithConfig(roleAdmin))
	onlineClassBookingDetails.DELETE("/:id", onlineClassController.DeleteOnlineClassBooking, md.CustomJWTWithConfig(roleAdmin))
	// online class category
	onlineClassCategory := onlineClasses.Group("/categories")
	onlineClassCategory.POST("", onlineClassController.CreateOnlineClassCategory, md.CustomJWTWithConfig(roleAdmin))
	onlineClassCategory.GET("", onlineClassController.GetOnlineClassCategories, md.CustomJWTWithConfig(roleUser))
	onlineClassCategoryDetails := onlineClassCategory.Group("/details")
	onlineClassCategoryDetails.GET("/:id", onlineClassController.GetOnlineClassCategoryDetail, md.CustomJWTWithConfig(allAccess))
	onlineClassCategoryDetails.PUT("/:id", onlineClassController.UpdateOnlineClassCategory, md.CustomJWTWithConfig(roleAdmin))
	onlineClassCategoryDetails.DELETE("/:id", onlineClassController.DeleteOnlineClassCategory, md.CustomJWTWithConfig(roleAdmin))

	// offline classes
	offlineClasses := v1.Group("/offline-classes")
	offlineClasses.POST("", offlineClassController.CreateOfflineClass, md.CustomJWTWithConfig(roleAdmin))
	offlineClasses.GET("", offlineClassController.GetOfflineClasses, md.CustomJWTWithConfig(allAccess))
	offlineClassDetails := offlineClasses.Group("/details")
	offlineClassDetails.GET("/:id", offlineClassController.GetOfflineClassDetail, md.CustomJWTWithConfig(allAccess))
	offlineClassDetails.PUT("/:id", offlineClassController.UpdateOfflineClass, md.CustomJWTWithConfig(roleAdmin))
	offlineClassDetails.DELETE("/:id", offlineClassController.DeleteOfflineClass, md.CustomJWTWithConfig(roleAdmin))
	// offline class booking
	offlineClassBooking := offlineClasses.Group("/bookings")
	offlineClassBooking.POST("", offlineClassController.CreateOfflineClassBooking, md.CustomJWTWithConfig(roleUser))
	offlineClassBooking.GET("", offlineClassController.GetOfflineClassBookings, md.CustomJWTWithConfig(roleAdmin))
	offlineClassBooking.GET("/user", offlineClassController.GetOfflineClassBookingUser, md.CustomJWTWithConfig(roleUser))
	offlineClassBooking.POST("/set-status/:id", offlineClassController.SetStatusOfflineClassBooking, md.CustomJWTWithConfig(roleAdmin))
	offlineClassBooking.POST("/pay/:id", offlineClassController.OfflineClassBookingPayment, md.CustomJWTWithConfig(roleUser))
	offlineClassBookingTake := offlineClassBooking.Group("/take")
	offlineClassBookingTake.GET("", offlineClassController.CheckOfflineClassBooking, md.CustomJWTWithConfig(roleAdmin))
	offlineClassBookingTake.POST("", offlineClassController.TakeOfflineClassBooking, md.CustomJWTWithConfig(roleAdmin))
	offlineClassBookingDetails := offlineClassBooking.Group("/details")
	offlineClassBookingDetails.GET("/:id", offlineClassController.GetOfflineClassBookingDetail, md.CustomJWTWithConfig(allAccess))
	offlineClassBookingDetails.PUT("/:id", offlineClassController.UpdateOfflineClassBooking, md.CustomJWTWithConfig(roleAdmin))
	offlineClassBookingDetails.DELETE("/:id", offlineClassController.DeleteOfflineClassBooking, md.CustomJWTWithConfig(roleAdmin))
	// offline class category
	offlineClassCategory := offlineClasses.Group("/categories")
	offlineClassCategory.POST("", offlineClassController.CreateOfflineClassCategory, md.CustomJWTWithConfig(roleAdmin))
	offlineClassCategory.GET("", offlineClassController.GetOfflineClassCategories, md.CustomJWTWithConfig(allAccess))
	offlineClassCategoryDetails := offlineClassCategory.Group("/details")
	offlineClassCategoryDetails.GET("/:id", offlineClassController.GetOfflineClassCategoryDetail, md.CustomJWTWithConfig(allAccess))
	offlineClassCategoryDetails.PUT("/:id", offlineClassController.UpdateOfflineClassCategory, md.CustomJWTWithConfig(roleAdmin))
	offlineClassCategoryDetails.DELETE("/:id", offlineClassController.DeleteOfflineClassCategory, md.CustomJWTWithConfig(roleAdmin))

	// history
	history := v1.Group("/histories")
	history.GET("/activities", historyController.FindHistoryActivity, md.CustomJWTWithConfig(roleUser))
	history.GET("/orders", historyController.FindHistoryOrder, md.CustomJWTWithConfig(roleUser))

	// create default user (superadmin)
	if err := userService.CreateSuperadmin(&model.User{
		Email:    config.Env.SUPERADMIN_EMAIL,
		Password: config.Env.SUPERADMIN_PASSWORD,
		Name:     config.Env.SUPERADMIN_NAME,
	}); err != nil {
		panic(err)
	}
}
