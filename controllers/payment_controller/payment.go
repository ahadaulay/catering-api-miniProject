package paymentcontroller

import (
	"catering-api/models/dto"
	paymentservice "catering-api/services/payment_service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PaymentController struct {
	PaymentService paymentservice.PaymentService
}

func (Ps *PaymentController) GetAllPayment(c echo.Context) error  {

	payment,err := Ps.PaymentService.GetAllPayment()

	if err != nil {
		c.JSON(http.StatusBadRequest,echo.Map{
			"message": "fail get all payment",
			"error":   err,
		})
	}

	return c.JSON(http.StatusAccepted,echo.Map{
		"message" : "success get all paymnet",
		"data" : payment,
	})
}

func (Pc *PaymentController) GetPaymentByID(c echo.Context) error   {
	var payment dto.PaymentResponse

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "Invalid payment ID",
        })
    }

	payment, err = Pc.PaymentService.GetPaymentByID(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get payment",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get payment by id",
		"data" : payment,
	})

}

func (Pc *PaymentController) CreatePayment(c echo.Context) error {
	var payment dto.PaymentCreate

	err := c.Bind(&payment)

	if err != nil {
		c.JSON(http.StatusBadRequest,echo.Map{
			"message" : "failed to bind data",
			"error" : err,
		})
	}

	err = Pc.PaymentService.CreatePayment(payment)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail create payment",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create payment",
	})
}

func (Pc *PaymentController) UpdatePayment(c echo.Context) error {
	var payment dto.PaymentResponse

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "Invalid payment ID",
        })
    }

	// Binding request body to struct
	if err := c.Bind(&payment); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request data",
			"error":   err.Error(),
		})
	}

	// Call service to update menu
	if err := Pc.PaymentService.UpdatePayment(id , payment); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to update payment",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Menu updated successfully",
	})
}


func(Pc *PaymentController) DeletePayment(c echo.Context) error  {

    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "Invalid payment ID",
        })
    }

	// Call service to delete course
	err = Pc.PaymentService.DeletePayment(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail delete payment",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete payment",
	})
}

