package paymentrepository

import "catering-api/models/dto"

type PaymentRepository interface {
	GetAllPayment() ([]dto.PaymentResponse, error)
	GetPaymentByID(id uint64) (dto.PaymentResponse, error)
	CreatePayment(input dto.PaymentCreate) error
	UpdatePayment(id uint64, input dto.PaymentResponse) error
	DeletePayment(id uint64) error
	GetPaymentByAdminID(adminID uint64) ([]dto.PaymentResponse, error)
}
