package admincontroller

import (
	"catering-api/helpers"
	"catering-api/helpers/middleware"
	"catering-api/helpers/response"
	"catering-api/models/dto"
	adminservice "catering-api/services/admin_service"
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type AdminController struct {
	AdminService adminservice.AdminService
}

func (Ac *AdminController) GetAllAdmin(c echo.Context) error {
	admin, err := Ac.AdminService.GetAllAdmin()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get all admin",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusAccepted, echo.Map{
		"message": "success get all admin",
		"data":    admin,
	})
}

func (Ac *AdminController) CreateAdmin(c echo.Context) error {
	var admin dto.AdminCreate

	err := c.Bind(&admin)

	if err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed to bind data",
			"error":   err,
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

func (Ac *AdminController) LoginAdmin(c echo.Context) error {
	AdminLoginRequest := dto.AdminLogin{}

	err := c.Bind(&AdminLoginRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("invalid client input"))
	}

	responses, err := Ac.AdminService.LoginAdmin(c, AdminLoginRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return c.JSON(http.StatusBadRequest, errors.New("invalid validation"))
		}

		if strings.Contains(err.Error(), "invalid email or password") {
			return c.JSON(http.StatusBadRequest, errors.New("invalid email or password"))
		}

		return c.JSON(http.StatusInternalServerError, errors.New("sign in error"))
	}

	userLoginResponse := response.AdminModelToAdminLoginResponse(responses)

	token, err := middleware.GenerateTokenAdmin(responses.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("generate jwt token error"))
	}

	userLoginResponse.Token = token

	return c.JSON(http.StatusCreated, helpers.SuccessResponse("Succesfully Sign In", userLoginResponse))
}
