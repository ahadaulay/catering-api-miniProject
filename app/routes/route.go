package routes

import (
	admincontroller "catering-api/controllers/admin_controller"
	foodcontroller "catering-api/controllers/food_controller"
	membershippackagecontroller "catering-api/controllers/membership_package_controller"
	membershiptransactioncontroller "catering-api/controllers/membership_transaction_controller"
	menucontroller "catering-api/controllers/menu_controller"
	menutransactioncontroller "catering-api/controllers/menu_transaction_controller"
	paymentcontroller "catering-api/controllers/payment_controller"
	usercontroller "catering-api/controllers/user_controller"
	"catering-api/helpers"
	"catering-api/helpers/middleware"
	adminrepository "catering-api/repositorys/admin_repository"
	foodrepository "catering-api/repositorys/food_repository"
	membershippackagerepository "catering-api/repositorys/membership_package_repository"
	membershiptransactionrepository "catering-api/repositorys/membership_transaction_repository"
	menurepository "catering-api/repositorys/menu_repository"
	menutransactionrepository "catering-api/repositorys/menu_transaction_repository"
	paymentrepository "catering-api/repositorys/payment_repository"
	userrepository "catering-api/repositorys/user_repository"
	adminservice "catering-api/services/admin_service"
	foodservice "catering-api/services/food_Service"
	membershippackageservice "catering-api/services/membership_package_service"
	membershiptransactionservice "catering-api/services/membership_transaction_service"
	menuservice "catering-api/services/menu_service"
	menutransactionservice "catering-api/services/menu_transaction_service"
	paymentservice "catering-api/services/payment_service"
	userservice "catering-api/services/user_service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	echoJwt "github.com/labstack/echo-jwt/v4"
)

func RouteService(db *gorm.DB) *echo.Echo {

	//route repository
	adminRouteRepository := adminrepository.NewAdminRepository(db)
	menuRouteRepository := menurepository.NewMenuRepository(db)
	membershipPackageRepository := membershippackagerepository.NewMenuRepository(db)
	paymentRepository := paymentrepository.NewPaymentRepository(db)
	userRepository := userrepository.NewUserRepository(db)
	foodRepository := foodrepository.NewFoodRepository(db)
	membershipTransactionRepository := membershiptransactionrepository.NewMenuRepository(db)
	menuTransactionRepository := menutransactionrepository.NewMenuTransactionRepository(db)

	//route service
	adminRouteService := adminservice.NewAdminService(adminRouteRepository)
	menuRouteService := menuservice.NewMenuService(menuRouteRepository)
	membershipPackageService := membershippackageservice.NewMenuService(membershipPackageRepository)
	paymentService := paymentservice.NewMenuService(paymentRepository)
	userService := userservice.NewMenuService(userRepository)
	foodService := foodservice.NewFoodService(foodRepository)
	membershipTransactionService := membershiptransactionservice.NewMenuService(membershipTransactionRepository)
	menuTransactionService := menutransactionservice.NewMenuService(menuTransactionRepository)

	//route controller
	AdminController := admincontroller.AdminController{
		AdminService: adminRouteService,
	}

	MenuController := menucontroller.MenuController{
		MenuService: menuRouteService,
	}

	MembershipPackgeController := membershippackagecontroller.MembershipPackageController{
		MembershipPackageService: membershipPackageService,
	}

	PaymentController := paymentcontroller.PaymentController{
		PaymentService: paymentService,
	}

	UserController := usercontroller.UserController{
		UserService: userService,
	}

	FoodController := foodcontroller.FoodController{
		FoodService: foodService,
	}

	membershipTransactionController := membershiptransactioncontroller.MembershipTransactionController{
		MembershipTransactionService: membershipTransactionService,
	}

	menuTransactionController := menutransactioncontroller.MenuTransactionController{
		MenuTransactionService: menuTransactionService,
	}

	app := echo.New()

	// ROUTES

	//midleware user
	userGroup := app.Group("user")
	userGroup.Use(echoJwt.JWT([]byte(helpers.GetConfig("USER_TOKEN_SECRET"))))

	//midleware admin
	AdminGroup := app.Group("admin")
	AdminGroup.Use(echoJwt.JWT([]byte(helpers.GetConfig("ADMIN_TOKEN_SECRET"))))

	//user routes
	AdminGroup.GET("/user", UserController.GetAllUser, middleware.AuthMiddleware("Admin"))
	app.GET("/user/:id", UserController.GetUserByID)
	app.POST("/user", UserController.CreateUser)
	userGroup.PUT("/user/:id", UserController.UpdateUser, middleware.AuthMiddleware("User"))
	AdminGroup.DELETE("/user/:id", UserController.DeleteUser, middleware.AuthMiddleware("Admin"))
	app.POST("/user/login", UserController.LoginUser)

	// Admin routes
	AdminGroup.GET("/admin", AdminController.GetAllAdmin, middleware.AuthMiddleware("Admin"))
	app.POST("/admin/login", AdminController.LoginAdmin)
	AdminGroup.GET("/admin/food/:admin_id", FoodController.GetFoodByAdminID, middleware.AuthMiddleware("Admin"))
	AdminGroup.GET("/admin/menu/:admin_id", MenuController.GetMenuByAdminID, middleware.AuthMiddleware("Admin"))
	AdminGroup.GET("/admin/payment/:admin_id", PaymentController.GetPaymentByAdminID, middleware.AuthMiddleware("Admin"))

	//Menu routes
	app.GET("/menu", MenuController.GetAllMenu)
	app.GET("/menu/:id", MenuController.GetMenuByID)
	AdminGroup.POST("/menu", MenuController.CreateMenu, middleware.AuthMiddleware("Admin"))
	AdminGroup.PUT("/menu/:id", MenuController.UpdateMenu, middleware.AuthMiddleware("Admin"))
	AdminGroup.DELETE("/menu/:id", MenuController.DeleteMenu, middleware.AuthMiddleware("Admin"))

	//food routes
	app.GET("/food", FoodController.GetAllFood)
	app.GET("/food/:id", FoodController.GetFoodByID)
	AdminGroup.POST("/food", FoodController.CreateFood, middleware.AuthMiddleware("Admin"))
	AdminGroup.PUT("/food/:id", FoodController.UpdateFood, middleware.AuthMiddleware("Admin"))
	AdminGroup.DELETE("/food/:id", FoodController.DeleteFood, middleware.AuthMiddleware("Admin"))

	//Membership Package routes
	app.GET("/membershippackage", MembershipPackgeController.GetAllMembershipPackage)
	app.GET("/membershippackage/:id", MembershipPackgeController.GetMembershipPackageByID)
	AdminGroup.POST("/membershippackage", MembershipPackgeController.CreateMembershipPackage, middleware.AuthMiddleware("Admin"))
	AdminGroup.PUT("/membershippackage/:id", MembershipPackgeController.UpdateMembershipPackage, middleware.AuthMiddleware("Admin"))
	AdminGroup.DELETE("/membershippackage/:id", MembershipPackgeController.DeleteMembershipPackage, middleware.AuthMiddleware("Admin"))

	//Membership transaction
	AdminGroup.GET("/membershiptransaction", membershipTransactionController.GetAllMembershipTransaction)
	AdminGroup.GET("/membershiptransaction/:id", membershipTransactionController.GetMembershipTransactionByID)
	userGroup.POST("/membershiptransaction", membershipTransactionController.CreateMembershipTransaction, middleware.AuthMiddleware("User"))
	AdminGroup.PUT("/membershiptransaction/:id", membershipTransactionController.UpdateMembershipTransaction, middleware.AuthMiddleware("Admin"))
	AdminGroup.DELETE("/membershiptransaction/:id", membershipTransactionController.DeleteMembershipTransaction, middleware.AuthMiddleware("Admin"))
	AdminGroup.PUT("/membershiptransaction/proof/:id", membershipTransactionController.UploadMembershipTransactionProof, middleware.AuthMiddleware("Admin"))
	AdminGroup.PUT("/membershiptransaction/accept/:id", membershipTransactionController.AcceptMembershipTransaction, middleware.AuthMiddleware("Admin"))
	AdminGroup.PUT("/membershiptransaction/reject/:id", membershipTransactionController.DeclineMembershipTransaction, middleware.AuthMiddleware("Admin"))

	//menu transaction route
	AdminGroup.GET("/menutransaction", menuTransactionController.GetAllMenuTransaction, middleware.AuthMiddleware("Admin"))
	AdminGroup.GET("/menutransaction/:id", menuTransactionController.GetMenuTransactionByID, middleware.AuthMiddleware("Admin"))
	userGroup.POST("/menutransaction", menuTransactionController.CreateMenuTransaction, middleware.AuthMiddleware("User"))
	AdminGroup.PUT("/menutransaction/:id", menuTransactionController.UpdateMenuTransaction, middleware.AuthMiddleware("Admin"))
	AdminGroup.DELETE("/menutransaction/:id", menuTransactionController.DeleteMenuTransaction, middleware.AuthMiddleware("Admin"))
	AdminGroup.PUT("/menutransaction/accept/:id", menuTransactionController.AcceptMenuTransaction, middleware.AuthMiddleware("Admin"))
	AdminGroup.PUT("/menutransaction/failed/:id", menuTransactionController.RejectMenuTransaction, middleware.AuthMiddleware("Admin"))
	AdminGroup.PUT("/menutransaction/success/:id", menuTransactionController.SuccessMenuTransaction, middleware.AuthMiddleware("Admin"))

	//Payment routes
	app.GET("/payment", PaymentController.GetAllPayment)
	app.GET("/payment/:id", PaymentController.GetPaymentByID)
	AdminGroup.POST("/payment", PaymentController.CreatePayment, middleware.AuthMiddleware("Admin"))
	AdminGroup.PUT("/payment/:id", PaymentController.UpdatePayment, middleware.AuthMiddleware("Admin"))
	AdminGroup.DELETE("payment/:id", PaymentController.DeletePayment, middleware.AuthMiddleware("Admin"))

	return app

}
