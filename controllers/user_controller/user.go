package usercontroller

import (
	"catering-api/helpers"
	"catering-api/helpers/middleware"
	"catering-api/helpers/response"
	"catering-api/models/dto"
	userservice "catering-api/services/user_service"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService userservice.UserService
}

func (Uc *UserController) GetAllUser(c echo.Context) error {

	user, err := Uc.UserService.GetAllUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{
			"message": "fail get all user",
			"error":   err,
		})
	}

	return c.JSON(http.StatusAccepted, echo.Map{
		"message": "success get all user",
		"data":    user,
	})
}

func (Uc *UserController) GetUserByID(c echo.Context) error {
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
		"data":    user,
	})

}

func (Uc *UserController) CreateUser(c echo.Context) error {
	var user dto.UserCreate

	err := c.Bind(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed to bind data",
			"error":   err,
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
	if err := Uc.UserService.UpdateUser(id, user); err != nil {
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

func (Uc *UserController) DeleteUser(c echo.Context) error {

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
	UserLoginRequest := dto.UserLogin{}

	err := c.Bind(&UserLoginRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request data",
			"error":   err.Error(),
		})
	}

	responses, err := Uc.UserService.LoginUser(c, UserLoginRequest)

	if err != nil {
		if strings.Contains(err.Error(), "invalid email or password") {
			return c.JSON(http.StatusBadRequest, errors.New("invalid email or password"))
		}

		return c.JSON(http.StatusInternalServerError, errors.New("sign in error"))
	}

	userLoginResponse := response.UserDomainToUserLoginResponse(responses)

	token, err := middleware.GenerateTokenUser(responses.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("generate jwt token error"))
	}

	userLoginResponse.Token = token

	return c.JSON(http.StatusCreated, helpers.SuccessResponse("Succesfully Sign In", userLoginResponse))
}
