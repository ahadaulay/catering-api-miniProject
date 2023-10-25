package foodcontroller

import (
	"catering-api/models/dto"
	foodservice "catering-api/services/food_Service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FoodController struct {
	FoodService foodservice.FoodService
}

func (Fc *FoodController) GetAllFood(c echo.Context) error  {

	food,err := Fc.FoodService.GetAllFood()

	if err != nil {
		c.JSON(http.StatusBadRequest,echo.Map{
			"message": "fail get all food",
			"error":   err,
		})
	}

	return c.JSON(http.StatusAccepted,echo.Map{
		"message" : "success get all food",
		"data" : food,
	})
}

func (Fc *FoodController) GetFoodByID(c echo.Context) error   {
	var food dto.FoodResponse

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "Invalid food ID",
        })
    }

	food, err = Fc.FoodService.GetFoodByID(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get menu",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get menu by id",
		"data" : food,
	})

}

func (Fc *FoodController) CreateFood(c echo.Context) error {
	var food dto.FoodCreate

	err := c.Bind(&food)

	if err != nil {
		c.JSON(http.StatusBadRequest,echo.Map{
			"message" : "failed to bind data",
			"error" : err,
		})
	}

	err = Fc.FoodService.CreateFood(food)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail create food",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create food",
	})
}

func (Fc *FoodController) UpdateFood(c echo.Context) error {
	var food dto.FoodCreate

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "Invalid food ID",
        })
    }

	// Binding request body to struct
	if err := c.Bind(&food); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request data",
			"error":   err.Error(),
		})
	}

	// Call service to update food
	if err := Fc.FoodService.UpdateFood(id , food); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to update food",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "food updated successfully",
	})
}


func(Fc *FoodController) DeleteFood(c echo.Context) error  {

    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "Invalid food ID",
        })
    }

	// Call service to delete course
	err = Fc.FoodService.DeleteFood(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail delete food",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete food",
	})
}

