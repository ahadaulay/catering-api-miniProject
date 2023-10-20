package routes

import (
	admincontroller "catering-api/controllers/admin_controller"
	adminrepository "catering-api/repositorys/admin_repository"
	adminservice "catering-api/services/admin_service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteService(db *gorm.DB) *echo.Echo {

	//route repository
	adminRouteRepository := adminrepository.NewAdminRepository(db)

	//route service
	adminRouteService := adminservice.NewAdminService(adminRouteRepository)


	//route controller
	AdminController := admincontroller.AdminController{
		AdminService : adminRouteService,
	}



	
	app := echo.New()

	// ROUTES
	// Admin routes
	app.GET("/admin",AdminController.GetAllAdmin)
	app.POST("/admin",AdminController.CreateAdmin)
	

	return app

}