package routes

import (
	"catering-api/app/middlewares"
	middlewareUser "catering-api/app/middlewares/user"
	admincontroller "catering-api/controllers/admin_controller"
	membershippackagecontroller "catering-api/controllers/membership_package_controller"
	menucontroller "catering-api/controllers/menu_controller"
	paymentcontroller "catering-api/controllers/payment_controller"
	usercontroller "catering-api/controllers/user_controller"
	"catering-api/helpers"
	adminrepository "catering-api/repositorys/admin_repository"
	membershippackagerepository "catering-api/repositorys/membership_package_repository"
	menurepository "catering-api/repositorys/menu_repository"
	paymentrepository "catering-api/repositorys/payment_repository"
	userrepository "catering-api/repositorys/user_repository"
	adminservice "catering-api/services/admin_service"
	membershippackageservice "catering-api/services/membership_package_service"
	menuservice "catering-api/services/menu_service"
	paymentservice "catering-api/services/payment_service"
	userservice "catering-api/services/user_service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func RouteService(db *gorm.DB) *echo.Echo {

	//route repository
	adminRouteRepository := adminrepository.NewAdminRepository(db)
	menuRouteRepository := menurepository.NewMenuRepository(db)
	membershipPackageRepository := membershippackagerepository.NewMenuRepository(db)
	paymentRepository := paymentrepository.NewPaymentRepository(db)
	userRepository := userrepository.NewUserRepository(db)

	//route service
	adminRouteService := adminservice.NewAdminService(adminRouteRepository)
	menuRouteService := menuservice.NewMenuService(menuRouteRepository)
	membershipPackageService := membershippackageservice.NewMenuService(membershipPackageRepository)
	paymentService := paymentservice.NewMenuService(paymentRepository)
	userService := userservice.NewMenuService(userRepository)


	//route controller
	AdminController := admincontroller.AdminController{
		AdminService : adminRouteService,
	}

	MenuController := menucontroller.MenuController{
		MenuService: menuRouteService,
	}

	MembershipPackgeController := membershippackagecontroller.MembershipPackageController{
		MembershipPackageService : membershipPackageService,
	}

	PaymentController := paymentcontroller.PaymentController{
		PaymentService: paymentService,
	}

	UserController := usercontroller.UserController{
		UserService: userService,
	}
	
	app := echo.New()

	configLogger := middlewares.ConfigLogger{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	configUser := middleware.JWTConfig{
		Claims: &middlewareUser.JwtUserClaims{},
		SigningKey: []byte(helpers.GetConfig("TOKEN_SECRET")),
	}

	app.Use(configLogger.Init())
	app.Use(middleware.CORS())



	
	

	// ROUTES
	//user routes
	app.GET("/user",UserController.GetAllUser)
	app.GET("/user:id",UserController.GetUserByID)
	app.POST("/user",UserController.CreateUser)
	app.PUT("/user/:id",UserController.UpdateUser)
	app.DELETE("/user/:id",UserController.DeleteUser)
	app.POST("/userLogin",UserController.LoginUser)

	privateUser := app.Group("/user", middleware.JWTWithConfig(configUser))
	privateUser.Use(middlewares.CheckTokenMiddlewareUser)
	privateUser.POST("/userLogout",UserController.LogoutUser)

	// Admin routes
	app.GET("/admin",AdminController.GetAllAdmin)
	app.POST("/admin",AdminController.CreateAdmin)

	//Menu routes
	privateUser.GET("/menu",MenuController.GetAllMenu)
	app.GET("/menu/:id",MenuController.GetMenuByID)
	app.POST("/menu",MenuController.CreateMenu)
	app.PUT("/menu/:id",MenuController.UpdateMenu)
	app.DELETE("/menu/:id",MenuController.DeleteMenu)

	//Membership Package routes
	app.GET("/membershippackage",MembershipPackgeController.GetAllMembershipPackage)
	app.GET("/membershippackage/:id",MembershipPackgeController.GetMembershipPackageByID)
	app.POST("/membershippackage",MembershipPackgeController.CreateMembershipPackage)
	app.PUT("/membershippackage/:id",MembershipPackgeController.UpdateMembershipPackage)
	app.DELETE("/membershippackage/:id",MembershipPackgeController.DeleteMembershipPackage)

	//Payment routes
	app.GET("/payment",PaymentController.GetAllPayment)
	app.GET("/payment/:id",PaymentController.GetPaymentByID)
	app.POST("/payment",PaymentController.CreatePayment)
	app.PUT("/payment/:id",PaymentController.UpdatePayment)
	app.DELETE("payment/:id",PaymentController.DeletePayment)
	

	return app

}