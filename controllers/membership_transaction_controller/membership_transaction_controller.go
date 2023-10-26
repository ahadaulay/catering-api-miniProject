package membershiptransactioncontroller

import (
	"catering-api/models/dto"
	membershiptransactionservice "catering-api/services/membership_transaction_service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MembershipTransactionController struct {
	MembershipTransactionService membershiptransactionservice.MembershipTransactionService
}

func (Mtc *MembershipTransactionController) GetAllMembershipTransaction(c echo.Context) error  {

	MembershipTransaction,err := Mtc.MembershipTransactionService.GetAllMembershipTransaction()

	if err != nil {
		c.JSON(http.StatusBadRequest,echo.Map{
			"message": "fail get all membership transaction",
			"error":   err,
		})
	}

	return c.JSON(http.StatusAccepted,echo.Map{
		"message" : "success get all membership transaction",
		"data" : MembershipTransaction,
	})
}

func (Mtc *MembershipTransactionController) GetMembershipTransactionByID(c echo.Context) error   {
	var MembershipTransaction dto.MembershipTransactionResponse

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "Invalid membership transaction ID",
        })
    }

	MembershipTransaction, err = Mtc.MembershipTransactionService.GetMembershipTransactionByID(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get membership transaction",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get membership transaction by id",
		"data" : MembershipTransaction,
	})

}

func (Mtc *MembershipTransactionController) CreateMembershipTransaction(c echo.Context) error {

	// // Dapatkan token JWT dari header Authorization
	// token := c.Request().Header.Get("Authorization")

	// // Periksa apakah token ada dan sesuai dengan format "Bearer [token]"
	// if token == "" {
	// 	return c.JSON(http.StatusUnauthorized, "Token tidak ada")
	// }

	var MembershipTransaction dto.MembershipTransactionCreate

	err := c.Bind(&MembershipTransaction)


	if err != nil {
		c.JSON(http.StatusBadRequest,echo.Map{
			"message" : "failed to bind data",
			"error" : err,
		})
	}

	err = Mtc.MembershipTransactionService.CreateMembershipTransaction(MembershipTransaction)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail create membership transaction",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create membership transaction",
	})
}

func (Mtc *MembershipTransactionController) UpdateMembershipTransaction(c echo.Context) error {
	var MembershipTransaction dto.MembershipTransactionCreate

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "Invalid membership transaction ID",
        })
    }

	// Binding request body to struct
	if err := c.Bind(&MembershipTransaction); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request data",
			"error":   err.Error(),
		})
	}

	// Call service to update membership transaction
	if err := Mtc.MembershipTransactionService.UpdateMembershipTransaction(id , MembershipTransaction); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to update membership transaction",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "membership transaction updated successfully",
	})
}


func(Mtc *MembershipTransactionController) DeleteMembershipTransaction(c echo.Context) error  {

    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "Invalid membership transaction ID",
			"id":id,
        })
    }

	// Call service to delete course
	err = Mtc.MembershipTransactionService.DeleteMembershipTransaction(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail delete membership transaction",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete membership transaction",
	})
}