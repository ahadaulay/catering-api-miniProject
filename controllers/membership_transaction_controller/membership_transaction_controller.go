package membershiptransactioncontroller

import (
	"catering-api/helpers"
	"catering-api/models/dto"
	membershiptransactionservice "catering-api/services/membership_transaction_service"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type MembershipTransactionController struct {
	MembershipTransactionService membershiptransactionservice.MembershipTransactionService
}

func (Mtc *MembershipTransactionController) GetAllMembershipTransaction(c echo.Context) error {

	MembershipTransaction, err := Mtc.MembershipTransactionService.GetAllMembershipTransaction()

	if err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{
			"message": "fail get all membership transaction",
			"error":   err,
		})
	}

	return c.JSON(http.StatusAccepted, echo.Map{
		"message": "success get all membership transaction",
		"data":    MembershipTransaction,
	})
}

func (Mtc *MembershipTransactionController) GetMembershipTransactionByID(c echo.Context) error {
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
		"data":    MembershipTransaction,
	})

}

func (Mtc *MembershipTransactionController) CreateMembershipTransaction(c echo.Context) error {

	var MembershipTransaction dto.MembershipTransactionCreate

	err := c.Bind(&MembershipTransaction)

	if err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed to bind data",
			"error":   err,
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
	if err := Mtc.MembershipTransactionService.UpdateMembershipTransaction(id, MembershipTransaction); err != nil {
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

func (Mtc *MembershipTransactionController) DeleteMembershipTransaction(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid membership transaction ID",
			"id":      id,
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

func (Mtc *MembershipTransactionController) UploadMembershipTransactionProof(c echo.Context) error {
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

	fileExtension := filepath.Ext(uniqueFilename)
	contentType := "application/octet-stream" // Nilai default jika ekstensi tidak dikenali

	// Daftar ekstensi gambar yang dikenali
	imageExtensions := map[string]string{
		".jpg":  "image/jpeg",
		".jpeg": "image/jpeg",
		".png":  "image/png",
		// Tambahkan ekstensi lain jika diperlukan
	}

	if val, ok := imageExtensions[fileExtension]; ok {
		contentType = val
	}

	params := &s3.PutObjectInput{
		Bucket: aws.String(helpers.GetConfig("AWS_BUCKET_NAME")),
		Key:    aws.String(uniqueFilename),
		Body:   src,
		// Gunakan tipe konten yang ditentukan oleh sistem (berdasarkan ekstensi)
		ContentType: aws.String(contentType),
		ACL:         aws.String("public-read"), // Set ACL ke public-read
	}

	_, err = svc.PutObject(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Dapatkan URL file yang diunggah
	imageURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", helpers.GetConfig("AWS_BUCKET_NAME"), uniqueFilename)

	MembershipTransaction.Proof = imageURL
	MembershipTransaction.Status = "waiting"

	// Call service to update membership transaction
	if err := Mtc.MembershipTransactionService.UpdateMembershipTransaction(id, MembershipTransaction); err != nil {
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

func (Mtc *MembershipTransactionController) AcceptMembershipTransaction(c echo.Context) error {
	var MembershipTransaction dto.MembershipTransactionResponse
	var MembershipTransactionCreate dto.MembershipTransactionCreate

	err := copier.Copy(&MembershipTransactionCreate, &MembershipTransaction)

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

	var user dto.UserCreate

	if err != nil {
		// Handle kesalahan jika diperlukan
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to copy user data",
		})
	}

	user, err = Mtc.MembershipTransactionService.GetUserByID(MembershipTransaction.UserID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get user",
			"error":   err.Error(),
		})
	}

	membershipPackage, err := Mtc.MembershipTransactionService.GetMembershipPackageByID(MembershipTransaction.MembershipPackageID)

	user.MembershipDuration += membershipPackage.Duration

	// Call service to update user
	if err := Mtc.MembershipTransactionService.UpdateUser(user.ID, user); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to update user",
			"error":   err.Error(),
		})
	}

	MembershipTransactionCreate.Status = "success"

	// Call service to update membership transaction
	if err := Mtc.MembershipTransactionService.UpdateMembershipTransaction(id, MembershipTransactionCreate); err != nil {
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

func (Mtc *MembershipTransactionController) DeclineMembershipTransaction(c echo.Context) error {
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
	MembershipTransaction.Status = "failed"

	// Call service to update membership transaction
	if err := Mtc.MembershipTransactionService.UpdateMembershipTransaction(id, MembershipTransaction); err != nil {
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
