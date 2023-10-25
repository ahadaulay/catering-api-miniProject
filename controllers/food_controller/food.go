package foodcontroller

import (
	"catering-api/helpers"
	"catering-api/models/dto"
	foodservice "catering-api/services/food_Service"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
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
			"message": "fail get food",
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
	svc, err := helpers.ConnectAWS()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }

    file, err := c.FormFile("image")
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }

    src, err := file.Open()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    defer src.Close()

    uniqueFilename := uuid.New().String() + "_" + time.Now().Format("20060102150405") + filepath.Ext(file.Filename)

    params := &s3.PutObjectInput{
        Bucket: aws.String(helpers.GetConfig("AWS_BUCKET_NAME")),
        Key:    aws.String(uniqueFilename),
        Body:   src,
    }

    _, err = svc.PutObject(params)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }

    // Dapatkan URL file yang diunggah
    imageURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", helpers.GetConfig("AWS_BUCKET_NAME"), uniqueFilename)

	var food dto.FoodCreate
	
	admin_id := c.FormValue("admin_id")
	admin_id_value, _ := strconv.ParseUint(admin_id, 10, 64)

	menu_id := c.FormValue("admin_id")
	menu_id_value, _ := strconv.ParseUint(menu_id, 10, 64)

	name := c.FormValue("name")

	food.Name = name
	food.AdminID = admin_id_value
	food.MenuID = menu_id_value
	food.Image = imageURL
	

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

