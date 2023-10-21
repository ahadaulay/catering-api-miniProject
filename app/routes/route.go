package routes

import (
	admincontroller "catering-api/controllers/admin_controller"
	menucontroller "catering-api/controllers/menu_controller"
	adminrepository "catering-api/repositorys/admin_repository"
	menurepository "catering-api/repositorys/menu_repository"
	adminservice "catering-api/services/admin_service"
	menuservice "catering-api/services/menu_service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteService(db *gorm.DB) *echo.Echo {

	//route repository
	adminRouteRepository := adminrepository.NewAdminRepository(db)
	menuRouteRepository := menurepository.NewMenuRepository(db)

	//route service
	adminRouteService := adminservice.NewAdminService(adminRouteRepository)
	menuRouteService := menuservice.NewMenuService(menuRouteRepository)


	//route controller
	AdminController := admincontroller.AdminController{
		AdminService : adminRouteService,
	}

	MenuController := menucontroller.MenuController{
		MenuService: menuRouteService,
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
	

	return app

}