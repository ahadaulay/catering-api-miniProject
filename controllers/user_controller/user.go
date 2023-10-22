package usercontroller

import (
	middlewares "catering-api/app/middlewares/user"
	"catering-api/models/dto"
	userservice "catering-api/services/user_service"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService userservice.UserService
}

func (Uc *UserController) GetAllUser(c echo.Context) error  {

	user,err := Uc.UserService.GetAllUser()

	if err != nil {
		c.JSON(http.StatusBadRequest,echo.Map{
			"message": "fail get all user",
			"error":   err,
		})
	}

	return c.JSON(http.StatusAccepted,echo.Map{
		"message" : "success get all user",
		"data" : user,
	})
}

func (Uc *UserController) GetUserByID(c echo.Context) error   {
	var user dto.UserResponse

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "Invalid user ID",
        })
    }

	user, err = Uc.UserService.GetUserByID(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get user",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get user by id",
		"data" : user,
	})

}

func (Uc *UserController) CreateUser(c echo.Context) error {
	var user dto.UserCreate

	err := c.Bind(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest,echo.Map{
			"message" : "failed to bind data",
			"error" : err,
		})
	}

	err = Uc.UserService.CreateUser(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail create user",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create user",
	})
}

func (Uc *UserController) UpdateUser(c echo.Context) error {
	var user dto.UserCreate

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "Invalid user ID",
        })
    }

	// Binding request body to struct
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request data",
			"error":   err.Error(),
		})
	}

	// Call service to update user
	if err := Uc.UserService.UpdateUser(id , user); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to update user",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "user updated successfully",
	})
}


func(Uc *UserController) DeleteUser(c echo.Context) error  {

    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "Invalid user id",
        })
    }

	// Call service to delete course
	err = Uc.UserService.DeleteUser(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail delete user",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete user",
	})
}

func (Uc *UserController) LoginUser(c echo.Context) error {
	var UserLogin dto.UserLogin
	err := c.Bind(&UserLogin)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail bind data",
			"error":   err.Error(),
		})
	}

	var user dto.UserResponse

	user, err = Uc.UserService.LoginUser(UserLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	token, errToken := middlewares.GenerateTokenUser(user.ID)

	if errToken != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail create token",
			"error":   errToken,
		})
	}

	userResponseLogin := dto.UserResponseLogin{
		ID: user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}

	return c.JSON(200, echo.Map{
		"message": "success login",
		"user":    userResponseLogin,
	})
}

func (Uc *UserController) LogoutUser(c echo.Context) error {
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


