package paymentservice

import (
	"catering-api/models/dto"
	paymentrepository "catering-api/repositorys/payment_repository"

	"github.com/jinzhu/copier"
)

type PaymentService interface {
	GetAllPayment() ([]dto.PaymentResponse, error)
	GetPaymentByID(uint64) (dto.PaymentResponse, error)
	CreatePayment(input dto.PaymentCreate) error
	UpdatePayment(id uint64 ,  input dto.PaymentResponse) error
	DeletePayment(id uint64) error
}

type PaymentImplementation struct {
	repository paymentrepository.PaymentRepository
}

func (Pi *PaymentImplementation) GetAllPayment() ([]dto.PaymentResponse, error) {
	data, err := Pi.repository.GetAllPayment()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (Pi *PaymentImplementation) GetPaymentByID(id uint64) (dto.PaymentResponse, error) {
	data, err := Pi.repository.GetPaymentByID(id)
	if err != nil {
		return dto.PaymentResponse{}, err
	}

	var Payment dto.PaymentResponse
	err = copier.Copy(&Payment, &data)
	if err != nil {
		return dto.PaymentResponse{}, err
	}

	return Payment, nil
}

func (Pi *PaymentImplementation) CreatePayment(input dto.PaymentCreate) error {
	err := Pi.repository.CreatePayment(input)
	if err != nil {
		return err
	}
	return nil
}

func (Pi *PaymentImplementation) UpdatePayment(id uint64, input dto.PaymentResponse) error {
	// call repository to update course
	err := Pi.repository.UpdatePayment(id,input)
	if err != nil {
		return err
	}
	return nil
}


func (Pi *PaymentImplementation) DeletePayment(id uint64) error {
	err := Pi.repository.DeletePayment(id)
	if err != nil {
		return err
	}
	return nil
}

func NewMenuService(paymentRepo paymentrepository.PaymentRepository) PaymentService {
	return &PaymentImplementation{
		repository: paymentRepo,
	}
}
