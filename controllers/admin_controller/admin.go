package admincontroller

import (
	middlewares "catering-api/app/middlewares/user"
	"catering-api/models/dto"
	adminservice "catering-api/services/admin_service"
	"net/http"

	"github.com/golang-jwt/jwt"
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

func (Ac *AdminController) LoginAdmin(c echo.Context) error {
	var AdminLogin dto.AdminLogin
	err := c.Bind(&AdminLogin)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail bind data",
			"error":   err.Error(),
		})
	}

	var admin dto.AdminResponse

	admin, err = Ac.AdminService.LoginAdmin(AdminLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	token, errToken := middlewares.GenerateTokenUser(admin.ID)

	if errToken != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail create token",
			"error":   errToken,
		})
	}

	adminResponseLogin := dto.AdminResponseLogin{
		ID: admin.ID,
		Name:  admin.Name,
		Email: admin.Email,
		Token: token,
	}

	return c.JSON(200, echo.Map{
		"message": "success login",
		"user":    adminResponseLogin,
	})
}

func (Ac *AdminController) LogoutUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)


	isListed := middlewares.CheckTokenUser(user.Raw)

	if !isListed {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "invalid token",
		})
	}

	middlewares.LogoutUser(user.Raw)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "logout success",
	})
}
