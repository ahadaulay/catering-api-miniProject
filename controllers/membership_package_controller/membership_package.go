package membershippackagecontroller

import (
	"catering-api/models/dto"
	membershippackageservice "catering-api/services/membership_package_service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MembershipPackageController struct {
	MembershipPackageService membershippackageservice.MembershipPackageService
}

func (Mpc *MembershipPackageController) GetAllMembershipPackage(c echo.Context) error {

	membershipPackage, err := Mpc.MembershipPackageService.GetAllMembershipPackage()

	if err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{
			"message": "fail get all membership package",
			"error":   err,
		})
	}

	return c.JSON(http.StatusAccepted, echo.Map{
		"message": "success get all membership package",
		"data":    membershipPackage,
	})
}

func (Mpc *MembershipPackageController) GetMembershipPackageByID(c echo.Context) error {
	var membershipPackage dto.MembershipPackageResponse

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid membership package ID",
		})
	}

	membershipPackage, err = Mpc.MembershipPackageService.GetMembershipPackageByID(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get membership package",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get menu by id",
		"data":    membershipPackage,
	})

}

func (Mpc *MembershipPackageController) CreateMembershipPackage(c echo.Context) error {
	var membershipPackage dto.MembershipPackageCreate

	err := c.Bind(&membershipPackage)

	if err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed to bind data",
			"error":   err,
		})
	}

	err = Mpc.MembershipPackageService.CreateMembershipPackage(membershipPackage)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail create membership package",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create membership package",
	})
}

func (Mpc *MembershipPackageController) UpdateMembershipPackage(c echo.Context) error {
	var membershipPackage dto.MembershipPackageResponse

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid membership package ID",
		})
	}

	// Binding request body to struct
	if err := c.Bind(&membershipPackage); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request membership package",
			"error":   err.Error(),
		})
	}

	// Call service to update menu
	if err := Mpc.MembershipPackageService.UpdateMembershipPackage(id, membershipPackage); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to update membership package",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "membership package updated successfully",
	})
}

func (Mpc *MembershipPackageController) DeleteMembershipPackage(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid membership package ID",
		})
	}

	// Call service to delete course
	err = Mpc.MembershipPackageService.DeleteMembershipPackage(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail delete membership package",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete membership package",
	})
}

func (Mpc *MembershipPackageController) GetMembershipPackageByAdminID(c echo.Context) error {
	var membershippackage []dto.MembershipPackageResponse

	admin_id, err := strconv.ParseUint(c.Param("admin_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid membership package ID",
		})
	}

	membershippackage, err = Mpc.MembershipPackageService.GetMembershipPackageByAdminID(admin_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get memebrship package",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get memebrship package by admin id",
		"data":    membershippackage,
	})

}
