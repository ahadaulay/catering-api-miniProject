package menucontroller

import (
	"catering-api/models/dto"
	menuservice "catering-api/services/menu_service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MenuController struct {
	MenuService menuservice.MenuService
}

func (Mc *MenuController) GetAllMenu(c echo.Context) error {

	menu, err := Mc.MenuService.GetAllMenu()

	if err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{
			"message": "fail get all menu",
			"error":   err,
		})
	}

	return c.JSON(http.StatusAccepted, echo.Map{
		"message": "success get all menu",
		"data":    menu,
	})
}

func (Mc *MenuController) GetMenuByID(c echo.Context) error {
	var menu dto.MenuResponse

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid menu ID",
		})
	}

	menu, err = Mc.MenuService.GetMenuByID(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get menu",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get menu by id",
		"data":    menu,
	})

}

func (Mc *MenuController) CreateMenu(c echo.Context) error {
	var menu dto.MenuCreate

	err := c.Bind(&menu)

	if err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed to bind data",
			"error":   err,
		})
	}

	err = Mc.MenuService.CreateMenu(menu)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail create menu",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create menu",
	})
}

func (Mc *MenuController) UpdateMenu(c echo.Context) error {
	var menu dto.MenuResponse

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid menu ID",
		})
	}

	// Binding request body to struct
	if err := c.Bind(&menu); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request data",
			"error":   err.Error(),
		})
	}

	// Call service to update menu
	if err := Mc.MenuService.UpdateMenu(id, menu); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to update menu",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Menu updated successfully",
	})
}

func (Mc *MenuController) DeleteMenu(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid menu ID",
		})
	}

	// Call service to delete course
	err = Mc.MenuService.DeleteMenu(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail delete menu",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete menu",
	})
}

func (Mc *MenuController) GetMenuByAdminID(c echo.Context) error {
	var menu []dto.MenuResponse

	admin_id, err := strconv.ParseUint(c.Param("admin_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid menu ID",
		})
	}

	menu, err = Mc.MenuService.GetMenuByAdminID(admin_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get menu",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get menu by admin id",
		"data":    menu,
	})

}
