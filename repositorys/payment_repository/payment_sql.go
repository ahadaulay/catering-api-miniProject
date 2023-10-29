package paymentrepository

import (
	"catering-api/models/dto"
	"catering-api/models/model"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type PaymentImplementation struct {
	db *gorm.DB
}

func (Pi *PaymentImplementation) GetAllPayment() ([]dto.PaymentResponse, error) {
	var PaymentModel []model.Payment

	err := Pi.db.Find(&PaymentModel).Error

	if err != nil {
		return nil, err
	}

	var Payment []dto.PaymentResponse

	if err := copier.Copy(&Payment, &PaymentModel); err != nil {
		return nil, err
	}

	return Payment, nil
}

func (Pi *PaymentImplementation) GetPaymentByID(id uint64) (dto.PaymentResponse, error) {
	var PaymentModel model.Payment

	err := Pi.db.Model(&model.Payment{}).Where("id = ? ", id).Find(&PaymentModel).Error

	if err != nil {
		return dto.PaymentResponse{}, err
	}

	// Periksa apakah data ditemukan
	if PaymentModel.ID == 0 {
		return dto.PaymentResponse{}, gorm.ErrRecordNotFound
	}

	var Payment dto.PaymentResponse

	err = copier.Copy(&Payment, &PaymentModel)

	if err != nil {
		return dto.PaymentResponse{}, err
	}

	return Payment, nil
}

func (Pi *PaymentImplementation) CreatePayment(input dto.PaymentCreate) error {
	var PaymentModel model.Payment

	err := copier.Copy(&PaymentModel, &input)

	if err != nil {
		return err
	}

	err = Pi.db.Model(&model.Payment{}).Create(&PaymentModel).Error

	if err != nil {
		return err
	}

	return nil
}

func (Pi *PaymentImplementation) UpdatePayment(id uint64, input dto.PaymentResponse) error {
	// Update menu with new data
	result := Pi.db.Model(&model.Payment{}).Where("id = ?", id).Updates(&model.Payment{
		AdminID:       input.AdminID,
		AccountName:   input.AccountName,
		AccountNumber: input.AccountName,
	})

	if result.Error != nil {
		return result.Error
	}

	// Check if any rows were affected
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (Pi *PaymentImplementation) DeletePayment(id uint64) error {
	err := Pi.db.Where("id = ?", id).Delete(&model.Payment{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (Pi *PaymentImplementation) GetPaymentByAdminID(adminID uint64) ([]dto.PaymentResponse, error) {
	var PaymentModel []model.Payment

	err := Pi.db.Model(&model.Payment{}).Where("admin_id = ? ", adminID).Find(&PaymentModel).Error

	// Periksa apakah FoodModel kosong (tidak ada data yang ditemukan)
	if len(PaymentModel) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	var payment []dto.PaymentResponse

	err = copier.Copy(&payment, &PaymentModel)

	if err != nil {
		return nil, err
	}

	return payment, nil
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &PaymentImplementation{
		db: db,
	}
}
