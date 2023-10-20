package admincontroller

import (
	"catering-api/models/dto"
	adminservice "catering-api/services/admin_service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminController struct {
	AdminService adminservice.AdminService
}

func (Ac *AdminController) GetAllAdmin(c echo.Context) error  {
	admin , err := Ac.AdminService.GetAllAdmin()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get all admin",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusAccepted,echo.Map{
		"message" : "success get all admin",
		"data" : admin,
	})
}

func(Ac *AdminController) CreateAdmin(c echo.Context) error  {
	var admin dto.AdminCreate

	err := c.Bind(&admin)

	if err != nil {
		c.JSON(http.StatusBadRequest,echo.Map{
			"message" : "failed to bind data",
			"error" : err,
		})
	}

	err = Ac.AdminService.CreateAdmin(admin)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail create admin",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create admin",
	})
}