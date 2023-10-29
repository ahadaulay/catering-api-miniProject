package menutransactioncontroller

import (
	"catering-api/models/dto"
	menutransactionservice "catering-api/services/menu_transaction_service"
	"github.com/jinzhu/copier"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MenuTransactionController struct {
	MenuTransactionService menutransactionservice.MenuTransactionService
}

func (Mtc *MenuTransactionController) GetAllMenuTransaction(c echo.Context) error {

	MenuTransaction, err := Mtc.MenuTransactionService.GetAllMenuTransaction()

	if err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{
			"message": "fail get all menu transaction",
			"error":   err,
		})
	}

	return c.JSON(http.StatusAccepted, echo.Map{
		"message": "success get all menu transaction",
		"data":    MenuTransaction,
	})
}

func (Mtc *MenuTransactionController) GetMenuTransactionByID(c echo.Context) error {
	var MenuTransaction dto.MenuTransactionResponse

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid menu transaction ID",
		})
	}

	MenuTransaction, err = Mtc.MenuTransactionService.GetMenuTransactionByID(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get menu transaction",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get menu transaction by id",
		"data":    MenuTransaction,
	})

}

func (Mtc *MenuTransactionController) CreateMenuTransaction(c echo.Context) error {
	// Deklarasikan MenuTransactionCreate
	var MenuTransaction dto.MenuTransactionCreate

	// Bind data dari request ke MenuTransaction
	if err := c.Bind(&MenuTransaction); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Failed to bind data",
			"error":   err.Error(),
		})
	}

	// Dapatkan pengguna berdasarkan UserID
	user, err := Mtc.MenuTransactionService.GetUserByID(MenuTransaction.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to get user information",
			"error":   err.Error(),
		})
	}

	// Periksa "membership duration"
	if user.MembershipDuration == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Your membership duration is 0, you are not allowed to place an order.",
			"error":   "Membership duration is 0",
		})
	}

	// Lanjutkan dengan membuat transaksi menu jika "membership duration" memungkinkan
	err = Mtc.MenuTransactionService.CreateMenuTransaction(MenuTransaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to create menu transaction",
			"error":   err.Error(),
		})
	}

	// Return response jika berhasil
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success create menu transaction",
	})
}

func (Mtc *MenuTransactionController) UpdateMenuTransaction(c echo.Context) error {
	var MenuTransaction dto.MenuTransactionCreate

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid menu transaction ID",
		})
	}

	// Binding request body to struct
	if err := c.Bind(&MenuTransaction); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request data",
			"error":   err.Error(),
		})
	}

	if err := Mtc.MenuTransactionService.UpdateMenuTransaction(id, MenuTransaction); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to update menu transaction",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "menu transaction updated successfully",
	})
}

func (Mtc *MenuTransactionController) DeleteMenuTransaction(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid menu transaction ID",
			"id":      id,
		})
	}

	// Call service to delete course
	err = Mtc.MenuTransactionService.DeleteMenuTransaction(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail delete menu transaction",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete menu transaction",
	})
}

func (Mtc *MenuTransactionController) AcceptMenuTransaction(c echo.Context) error {
	var MenuTransaction dto.MenuTransactionCreate
	var GetMenuTransaction dto.MenuTransactionResponse

	err := copier.Copy(&MenuTransaction, &GetMenuTransaction)

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid menu transaction ID",
		})
	}

	GetMenuTransaction, err = Mtc.MenuTransactionService.GetMenuTransactionByID(id)

	id, err = strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid menu transaction ID",
		})
	}

	// Binding request body to struct
	if err := c.Bind(&MenuTransaction); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request data",
			"error":   err.Error(),
		})
	}

	MenuTransaction.Status = "accept"

	if err := Mtc.MenuTransactionService.ReduceMenuStock(GetMenuTransaction.MenuID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to make transactionz",
			"error":   err.Error(),
		})

	}

	if err := Mtc.MenuTransactionService.ReduceUserDuration(GetMenuTransaction.UserID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to make transactionz",
			"error":   err.Error(),
		})

	}

	if err := Mtc.MenuTransactionService.UpdateMenuTransaction(id, MenuTransaction); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to accept menu transaction",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "menu transaction accept successfully",
	})
}

func (Mtc *MenuTransactionController) RejectMenuTransaction(c echo.Context) error {
	var MenuTransaction dto.MenuTransactionCreate

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid menu transaction ID",
		})
	}

	// Binding request body to struct
	if err := c.Bind(&MenuTransaction); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request data",
			"error":   err.Error(),
		})
	}

	MenuTransaction.Status = "failed"

	if err := Mtc.MenuTransactionService.UpdateMenuTransaction(id, MenuTransaction); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to update menu transaction",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "menu transaction updated successfully",
	})
}

func (Mtc *MenuTransactionController) SuccessMenuTransaction(c echo.Context) error {
	var MenuTransaction dto.MenuTransactionCreate

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid menu transaction ID",
		})
	}

	// Binding request body to struct
	if err := c.Bind(&MenuTransaction); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request data",
			"error":   err.Error(),
		})
	}

	MenuTransaction.Status = "success"

	if err := Mtc.MenuTransactionService.UpdateMenuTransaction(id, MenuTransaction); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to update menu transaction",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "menu transaction updated successfully",
	})
}
