package route

import (
	"github.com/Group10CapstoneProject/Golang/config"
	"github.com/Group10CapstoneProject/Golang/constans"
	pkgArticleController "github.com/Group10CapstoneProject/Golang/internal/articles/controller"
	pkgArticleRepostiory "github.com/Group10CapstoneProject/Golang/internal/articles/repository"
	pkgArticleService "github.com/Group10CapstoneProject/Golang/internal/articles/service"
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
	pkgTrainerController "github.com/Group10CapstoneProject/Golang/internal/trainers/controller"
	pkgTrainerRepostiory "github.com/Group10CapstoneProject/Golang/internal/trainers/repository"
	pkgTrainerService "github.com/Group10CapstoneProject/Golang/internal/trainers/service"
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
	trainerRepository := pkgTrainerRepostiory.NewTrainerRepository(db)
	articleRepository := pkgArticleRepostiory.NewArticlesRepository(db)

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
	historyService := pkgHistoryService.NewHistoryService(memberRepository, onlineClassRepository, offlineClassRepository, trainerRepository)
	trainerService := pkgTrainerService.NewTrainerService(trainerRepository, memberRepository, notificationRepository, imagekitService)
	articleService := pkgArticleService.NewArticlesService(articleRepository)

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
	trainerController := pkgTrainerController.NewTrainerController(memberService, jwtService, noticationService, trainerService)
	articleController := pkgArticleController.NewArticlesController(articleService)

	// int route
	// auth
	auth := v1.Group("/auth")
	auth.POST("/login", authController.Login).Name = "login"
	auth.POST("/refresh", authController.RefreshToken).Name = "refresh-token-user"
	auth.POST("/logout", authController.Logout, md.CustomJWTWithConfig(allAccess)).Name = "logout"
	admin := auth.Group("/admin")
	admin.POST("/login", authController.LoginAdmin).Name = "login-admin"
	admin.POST("/refresh", authController.RefreshAdminToken).Name = "refresh-token-admin"

	// users
	users := v1.Group("/users")
	users.GET("", userController.GetUsers, md.CustomJWTWithConfig(roleAdmin)).Name = "get-all-users"
	users.POST("/signup", userController.Signup).Name = "signup-user"
	users.GET("/profile", userController.GetUser, md.CustomJWTWithConfig(allAccess)).Name = "get-user-profile"
	userAdmin := users.Group("/admin")
	userAdmin.POST("", userController.NewAadmin, md.CustomJWTWithConfig(roleSuperadmin)).Name = "create-admin"
	userAdmin.GET("", userController.GetAdmins, md.CustomJWTWithConfig(roleSuperadmin)).Name = "get-all-admins"

	// payment methods
	paymentMethods := v1.Group("/payment-methods")
	paymentMethods.POST("", paymentMethodController.CreatePaymentMethod, md.CustomJWTWithConfig(roleAdmin)).Name = "create-payment-method"
	paymentMethods.GET("", paymentMethodController.GetPaymentMethods, md.CustomJWTWithConfig(allAccess)).Name = "get-all-payment-methods"
	paymentMethodDetails := paymentMethods.Group("/details")
	paymentMethodDetails.GET("/:id", paymentMethodController.GetPaymentMethodDetail, md.CustomJWTWithConfig(allAccess)).Name = "get-payment-method-detail"
	paymentMethodDetails.PUT("/:id", paymentMethodController.UpdatePaymentMethod, md.CustomJWTWithConfig(roleAdmin)).Name = "update-payment-method"
	paymentMethodDetails.DELETE("/:id", paymentMethodController.DeletePaymentMethod, md.CustomJWTWithConfig(roleAdmin)).Name = "delete-payment-method"

	// members
	members := v1.Group("/members")
	members.POST("", memberController.CreateMember, md.CustomJWTWithConfig(roleUser)).Name = "create-member"
	members.GET("", memberController.GetMembers, md.CustomJWTWithConfig(roleAdmin)).Name = "get-all-members"
	members.POST("/set-status/:id", memberController.SetStatusMember, md.CustomJWTWithConfig(roleAdmin)).Name = "set-status-member"
	members.POST("/pay/:id", memberController.MemberPayment, md.CustomJWTWithConfig(roleUser)).Name = "member-payment"
	members.GET("/user", memberController.GetMemberUser, md.CustomJWTWithConfig(roleUser)).Name = "get-member-user"
	memberDetails := members.Group("/details")
	memberDetails.GET("/:id", memberController.GetMemberDetail, md.CustomJWTWithConfig(allAccess)).Name = "get-member-detail"
	memberDetails.PUT("/:id", memberController.UpdateMember, md.CustomJWTWithConfig(roleAdmin)).Name = "update-member"
	memberDetails.DELETE("/:id", memberController.DeleteMember, md.CustomJWTWithConfig(roleAdmin)).Name = "delete-member"
	// member types
	memberTypes := members.Group("/types")
	memberTypes.POST("", memberController.CreateMemberType, md.CustomJWTWithConfig(roleAdmin)).Name = "create-member-type"
	memberTypes.GET("", memberController.GetMemberTypes, md.CustomJWTWithConfig(allAccess)).Name = "get-all-member-types"
	memberTypeDetails := memberTypes.Group("/details")
	memberTypeDetails.GET("/:id", memberController.GetMemberTypeDetail, md.CustomJWTWithConfig(allAccess)).Name = "get-member-type-detail"
	memberTypeDetails.PUT("/:id", memberController.UpdateMemberType, md.CustomJWTWithConfig(roleAdmin)).Name = "update-member-type"
	memberTypeDetails.DELETE("/:id", memberController.DeleteMemberType, md.CustomJWTWithConfig(roleAdmin)).Name = "delete-member-type"

	// files
	files := v1.Group("/files")
	files.POST("/upload", fileController.Upload, md.CustomJWTWithConfig(roleAdmin)).Name = "upload-file"

	// notifications
	notifications := v1.Group("/notifications")
	notifications.GET("", noticationController.GetNotifications, md.CustomJWTWithConfig(roleAdmin)).Name = "get-all-notifications"

	// online classes
	onlineClasses := v1.Group("/online-classes")
	onlineClasses.POST("", onlineClassController.CreateOnlineClass, md.CustomJWTWithConfig(roleAdmin)).Name = "create-online-class"
	onlineClasses.GET("", onlineClassController.GetOnlineClasses, md.CustomJWTWithConfig(allAccess)).Name = "get-all-online-classes"
	onlineClassDetails := onlineClasses.Group("/details")
	onlineClassDetails.GET("/:id", onlineClassController.GetOnlineClassDetail, md.CustomJWTWithConfig(allAccess)).Name = "get-online-class-detail"
	onlineClassDetails.PUT("/:id", onlineClassController.UpdateOnlineClass, md.CustomJWTWithConfig(roleAdmin)).Name = "update-online-class"
	onlineClassDetails.DELETE("/:id", onlineClassController.DeleteOnlineClass, md.CustomJWTWithConfig(roleAdmin)).Name = "delete-online-class"
	// online class booking
	onlineClassBooking := onlineClasses.Group("/bookings")
	onlineClassBooking.POST("", onlineClassController.CreateOnlineClassBooking, md.CustomJWTWithConfig(roleUser)).Name = "create-online-class-booking"
	onlineClassBooking.GET("", onlineClassController.GetOnlineClassBookings, md.CustomJWTWithConfig(roleAdmin)).Name = "get-all-online-class-bookings"
	onlineClassBooking.GET("/user", onlineClassController.GetOnlineClassBookingUser, md.CustomJWTWithConfig(roleUser)).Name = "get-online-class-booking-user"
	onlineClassBooking.POST("/set-status/:id", onlineClassController.SetStatusOnlineClassBooking, md.CustomJWTWithConfig(roleAdmin)).Name = "set-status-online-class-booking"
	onlineClassBooking.POST("/pay/:id", onlineClassController.OnlineClassBookingPayment, md.CustomJWTWithConfig(roleUser)).Name = "online-class-booking-payment"
	onlineClassBookingDetails := onlineClassBooking.Group("/details")
	onlineClassBookingDetails.GET("/:id", onlineClassController.GetOnlineClassBookingDetail, md.CustomJWTWithConfig(allAccess)).Name = "get-online-class-booking-detail"
	onlineClassBookingDetails.PUT("/:id", onlineClassController.UpdateOnlineClassBooking, md.CustomJWTWithConfig(roleAdmin)).Name = "update-online-class-booking"
	onlineClassBookingDetails.DELETE("/:id", onlineClassController.DeleteOnlineClassBooking, md.CustomJWTWithConfig(roleAdmin)).Name = "delete-online-class-booking"
	// online class category
	onlineClassCategory := onlineClasses.Group("/categories")
	onlineClassCategory.POST("", onlineClassController.CreateOnlineClassCategory, md.CustomJWTWithConfig(roleAdmin)).Name = "create-online-class-category"
	onlineClassCategory.GET("", onlineClassController.GetOnlineClassCategories, md.CustomJWTWithConfig(roleUser)).Name = "get-all-online-class-categories"
	onlineClassCategoryDetails := onlineClassCategory.Group("/details")
	onlineClassCategoryDetails.GET("/:id", onlineClassController.GetOnlineClassCategoryDetail, md.CustomJWTWithConfig(allAccess)).Name = "get-online-class-category-detail"
	onlineClassCategoryDetails.PUT("/:id", onlineClassController.UpdateOnlineClassCategory, md.CustomJWTWithConfig(roleAdmin)).Name = "update-online-class-category"
	onlineClassCategoryDetails.DELETE("/:id", onlineClassController.DeleteOnlineClassCategory, md.CustomJWTWithConfig(roleAdmin)).Name = "delete-online-class-category"

	// offline classes
	offlineClasses := v1.Group("/offline-classes")
	offlineClasses.POST("", offlineClassController.CreateOfflineClass, md.CustomJWTWithConfig(roleAdmin)).Name = "create-offline-class"
	offlineClasses.GET("", offlineClassController.GetOfflineClasses, md.CustomJWTWithConfig(allAccess)).Name = "get-all-offline-classes"
	offlineClassDetails := offlineClasses.Group("/details")
	offlineClassDetails.GET("/:id", offlineClassController.GetOfflineClassDetail, md.CustomJWTWithConfig(allAccess)).Name = "get-offline-class-detail"
	offlineClassDetails.PUT("/:id", offlineClassController.UpdateOfflineClass, md.CustomJWTWithConfig(roleAdmin)).Name = "update-offline-class"
	offlineClassDetails.DELETE("/:id", offlineClassController.DeleteOfflineClass, md.CustomJWTWithConfig(roleAdmin)).Name = "delete-offline-class"
	// offline class booking
	offlineClassBooking := offlineClasses.Group("/bookings")
	offlineClassBooking.POST("", offlineClassController.CreateOfflineClassBooking, md.CustomJWTWithConfig(roleUser)).Name = "create-offline-class-booking"
	offlineClassBooking.GET("", offlineClassController.GetOfflineClassBookings, md.CustomJWTWithConfig(roleAdmin)).Name = "get-all-offline-class-bookings"
	offlineClassBooking.GET("/user", offlineClassController.GetOfflineClassBookingUser, md.CustomJWTWithConfig(roleUser)).Name = "get-offline-class-booking-user"
	offlineClassBooking.POST("/set-status/:id", offlineClassController.SetStatusOfflineClassBooking, md.CustomJWTWithConfig(roleAdmin)).Name = "set-status-offline-class-booking"
	offlineClassBooking.POST("/pay/:id", offlineClassController.OfflineClassBookingPayment, md.CustomJWTWithConfig(roleUser)).Name = "offline-class-booking-payment"
	offlineClassBookingTake := offlineClassBooking.Group("/take")
	offlineClassBookingTake.GET("", offlineClassController.CheckOfflineClassBooking, md.CustomJWTWithConfig(roleAdmin)).Name = "check-offline-class-booking"
	offlineClassBookingTake.POST("", offlineClassController.TakeOfflineClassBooking, md.CustomJWTWithConfig(roleAdmin)).Name = "take-offline-class-booking"
	offlineClassBookingDetails := offlineClassBooking.Group("/details")
	offlineClassBookingDetails.GET("/:id", offlineClassController.GetOfflineClassBookingDetail, md.CustomJWTWithConfig(allAccess)).Name = "get-offline-class-booking-detail"
	offlineClassBookingDetails.PUT("/:id", offlineClassController.UpdateOfflineClassBooking, md.CustomJWTWithConfig(roleAdmin)).Name = "update-offline-class-booking"
	offlineClassBookingDetails.DELETE("/:id", offlineClassController.DeleteOfflineClassBooking, md.CustomJWTWithConfig(roleAdmin)).Name = "delete-offline-class-booking"
	// offline class category
	offlineClassCategory := offlineClasses.Group("/categories")
	offlineClassCategory.POST("", offlineClassController.CreateOfflineClassCategory, md.CustomJWTWithConfig(roleAdmin)).Name = "create-offline-class-category"
	offlineClassCategory.GET("", offlineClassController.GetOfflineClassCategories, md.CustomJWTWithConfig(allAccess)).Name = "get-all-offline-class-categories"
	offlineClassCategoryDetails := offlineClassCategory.Group("/details")
	offlineClassCategoryDetails.GET("/:id", offlineClassController.GetOfflineClassCategoryDetail, md.CustomJWTWithConfig(allAccess)).Name = "get-offline-class-category-detail"
	offlineClassCategoryDetails.PUT("/:id", offlineClassController.UpdateOfflineClassCategory, md.CustomJWTWithConfig(roleAdmin)).Name = "update-offline-class-category"
	offlineClassCategoryDetails.DELETE("/:id", offlineClassController.DeleteOfflineClassCategory, md.CustomJWTWithConfig(roleAdmin)).Name = "delete-offline-class-category"

	// history
	history := v1.Group("/histories")
	history.GET("/activities", historyController.FindHistoryActivity, md.CustomJWTWithConfig(roleUser)).Name = "get-history-activity"
	history.GET("/orders", historyController.FindHistoryOrder, md.CustomJWTWithConfig(roleUser)).Name = "get-history-order"

	// trainer
	trainer := v1.Group("/trainers")
	trainer.POST("", trainerController.CreateTrainer, md.CustomJWTWithConfig(roleAdmin)).Name = "create-trainer"
	trainer.GET("", trainerController.GetTrainers, md.CustomJWTWithConfig(allAccess)).Name = "get-all-trainers"
	trainerDetails := trainer.Group("/details")
	trainerDetails.GET("/:id", trainerController.GetTrainerDetail, md.CustomJWTWithConfig(allAccess)).Name = "get-trainer-detail"
	trainerDetails.PUT("/:id", trainerController.UpdateTrainer, md.CustomJWTWithConfig(roleAdmin)).Name = "update-trainer"
	trainerDetails.DELETE("/:id", trainerController.DeleteTrainer, md.CustomJWTWithConfig(roleAdmin)).Name = "delete-trainer"
	trainerSkill := trainer.Group("/skills")
	trainerSkill.POST("", trainerController.CreateSkill, md.CustomJWTWithConfig(roleAdmin)).Name = "create-skill"
	trainerSkill.GET("", trainerController.GetSkills, md.CustomJWTWithConfig(allAccess)).Name = "get-all-skills"
	trainerSkillDetails := trainerSkill.Group("/details")
	trainerSkillDetails.GET("/:id", trainerController.GetSkillDetail, md.CustomJWTWithConfig(allAccess)).Name = "get-skill-detail"
	trainerSkillDetails.PUT("/:id", trainerController.UpdateSkill, md.CustomJWTWithConfig(roleAdmin)).Name = "update-skill"
	trainerSkillDetails.DELETE("/:id", trainerController.DeleteSkill, md.CustomJWTWithConfig(roleAdmin)).Name = "delete-skill"
	trainerBooking := trainer.Group("/bookings")
	trainerBooking.POST("", trainerController.CreateTrainerBooking, md.CustomJWTWithConfig(roleUser)).Name = "create-trainer-booking"
	trainerBooking.GET("", trainerController.GetTrainerBookings, md.CustomJWTWithConfig(roleAdmin)).Name = "get-all-trainer-bookings"
	trainerBooking.GET("/user", trainerController.GetTrainerBookingUser, md.CustomJWTWithConfig(roleUser)).Name = "get-trainer-booking-user"
	trainerBooking.POST("/set-status/:id", trainerController.SetStatusTrainerBooking, md.CustomJWTWithConfig(roleAdmin)).Name = "set-status-trainer-booking"
	trainerBooking.POST("/pay/:id", trainerController.TrainerBookingPayment, md.CustomJWTWithConfig(roleUser)).Name = "trainer-booking-payment"
	trainerBookingTake := trainerBooking.Group("/take")
	trainerBookingTake.GET("", trainerController.CheckTrainerBooking, md.CustomJWTWithConfig(roleAdmin)).Name = "check-trainer-booking"
	trainerBookingTake.POST("", trainerController.TakeTrainerBooking, md.CustomJWTWithConfig(roleAdmin)).Name = "take-trainer-booking"
	trainerBookingDetails := trainerBooking.Group("/details")
	trainerBookingDetails.GET("/:id", trainerController.GetTrainerBookingDetail, md.CustomJWTWithConfig(allAccess)).Name = "get-trainer-booking-detail"
	trainerBookingDetails.PUT("/:id", trainerController.UpdateTrainerBooking, md.CustomJWTWithConfig(roleAdmin)).Name = "update-trainer-booking"
	trainerBookingDetails.DELETE("/:id", trainerController.DeleteTrainerBooking, md.CustomJWTWithConfig(roleAdmin)).Name = "delete-trainer-booking"

	// articles
	article := v1.Group("/articles")
	article.POST("", articleController.CreateArticles, md.CustomJWTWithConfig(roleAdmin)).Name = "create-article"
	article.GET("", articleController.GetArticles, md.CustomJWTWithConfig(allAccess)).Name = "get-all-articles"
	articleDetails := article.Group("/details")
	articleDetails.GET("/:id", articleController.GetArticlesDetail, md.CustomJWTWithConfig(allAccess)).Name = "get-article-detail"
	articleDetails.PUT("/:id", articleController.UpdateArticles, md.CustomJWTWithConfig(roleAdmin)).Name = "update-article"
	articleDetails.DELETE("/:id", articleController.DeleteArticles, md.CustomJWTWithConfig(roleAdmin)).Name = "delete-article"

	// create default user (superadmin)
	if err := userService.CreateSuperadmin(&model.User{
		Email:    config.Env.SUPERADMIN_EMAIL,
		Password: config.Env.SUPERADMIN_PASSWORD,
		Name:     config.Env.SUPERADMIN_NAME,
	}); err != nil {
		panic(err)
	}
}
