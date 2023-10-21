package routes

import (
	admincontroller "catering-api/controllers/admin_controller"
	membershippackagecontroller "catering-api/controllers/membership_package_controller"
	menucontroller "catering-api/controllers/menu_controller"
	paymentcontroller "catering-api/controllers/payment_controller"
	adminrepository "catering-api/repositorys/admin_repository"
	membershippackagerepository "catering-api/repositorys/membership_package_repository"
	menurepository "catering-api/repositorys/menu_repository"
	paymentrepository "catering-api/repositorys/payment_repository"
	adminservice "catering-api/services/admin_service"
	membershippackageservice "catering-api/services/membership_package_service"
	menuservice "catering-api/services/menu_service"
	paymentservice "catering-api/services/payment_service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteService(db *gorm.DB) *echo.Echo {

	//route repository
	adminRouteRepository := adminrepository.NewAdminRepository(db)
	menuRouteRepository := menurepository.NewMenuRepository(db)
	membershipPackageRepository := membershippackagerepository.NewMenuRepository(db)
	paymentRepository := paymentrepository.NewPaymentRepository(db)

	//route service
	adminRouteService := adminservice.NewAdminService(adminRouteRepository)
	menuRouteService := menuservice.NewMenuService(menuRouteRepository)
	membershipPackageService := membershippackageservice.NewMenuService(membershipPackageRepository)
	paymentService := paymentservice.NewMenuService(paymentRepository)


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



	
	app := echo.New()

	// ROUTES
	// Admin routes
	app.GET("/admin",AdminController.GetAllAdmin)
	app.POST("/admin",AdminController.CreateAdmin)

	//Menu routes
	app.GET("/menu",MenuController.GetAllMenu)
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