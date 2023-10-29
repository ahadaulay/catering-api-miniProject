package menutransactionservice

import (
	"catering-api/models/dto"
	menutransactionrepository "catering-api/repositorys/menu_transaction_repository"

	"github.com/jinzhu/copier"
)

type MenuTransactionService interface {
	GetAllMenuTransaction() ([]dto.MenuTransactionResponse, error)
	GetMenuTransactionByID(id uint64) (dto.MenuTransactionResponse, error)
	CreateMenuTransaction(input dto.MenuTransactionCreate) error
	UpdateMenuTransaction(id uint64, input dto.MenuTransactionCreate) error
	DeleteMenuTransaction(id uint64) error
	ReduceMenuStock(id uint64) error
	ReduceUserDuration(id uint64) error
	GetUserByID(id uint64) (dto.UserResponse, error)
}

type MenuTransactionImplementation struct {
	repository menutransactionrepository.MenuTransactionRepository
}

func (Mti *MenuTransactionImplementation) GetAllMenuTransaction() ([]dto.MenuTransactionResponse, error) {
	data, err := Mti.repository.GetAllMenuTransaction()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (Mti *MenuTransactionImplementation) GetMenuTransactionByID(id uint64) (dto.MenuTransactionResponse, error) {
	data, err := Mti.repository.GetMenuTransactionByID(id)
	if err != nil {
		return dto.MenuTransactionResponse{}, err
	}

	var menushiptransaction dto.MenuTransactionResponse
	err = copier.Copy(&menushiptransaction, &data)
	if err != nil {
		return dto.MenuTransactionResponse{}, err
	}

	return menushiptransaction, nil
}

func (Mti *MenuTransactionImplementation) CreateMenuTransaction(input dto.MenuTransactionCreate) error {
	err := Mti.repository.CreateMenuTransaction(input)
	if err != nil {
		return err
	}
	return nil
}

func (Mti *MenuTransactionImplementation) UpdateMenuTransaction(id uint64, input dto.MenuTransactionCreate) error {
	// call repository to update course
	err := Mti.repository.UpdateMenuTransaction(id, input)
	if err != nil {
		return err
	}
	return nil
}

func (Mti *MenuTransactionImplementation) DeleteMenuTransaction(id uint64) error {
	err := Mti.repository.DeleteMenuTransaction(id)
	if err != nil {
		return err
	}
	return nil
}

func (Mti *MenuTransactionImplementation) GetUserByID(id uint64) (dto.UserResponse, error) {
	data, err := Mti.repository.GetUserByID(id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	var user dto.UserResponse
	err = copier.Copy(&user, &data)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return user, nil
}

func (Mti *MenuTransactionImplementation) ReduceMenuStock(id uint64) error {

	err := Mti.repository.ReduceMenuStock(id)
	if err != nil {
		return err
	}
	return nil
}

func (Mti *MenuTransactionImplementation) ReduceUserDuration(id uint64) error {

	err := Mti.repository.ReduceDurationUser(id)
	if err != nil {
		return err
	}
	return nil
}

func NewMenuService(menuTransactionRepo menutransactionrepository.MenuTransactionRepository) MenuTransactionService {
	return &MenuTransactionImplementation{
		repository: menuTransactionRepo,
	}
}
