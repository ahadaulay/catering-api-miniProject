package membershiptransactionservice

import (
	"catering-api/helpers"
	"catering-api/models/dto"
	membershiptransactionrepository "catering-api/repositorys/membership_transaction_repository"

	"github.com/jinzhu/copier"
)

type MembershipTransactionService interface {
	GetAllMembershipTransaction() ([]dto.MembershipTransactionResponse, error)
	GetMembershipTransactionByID(id uint64) (dto.MembershipTransactionResponse, error)
	CreateMembershipTransaction(input dto.MembershipTransactionCreate) error
	UpdateMembershipTransaction(id uint64, input dto.MembershipTransactionCreate) error
	DeleteMembershipTransaction(id uint64) error
	GetUserByID(id uint64) (dto.UserCreate, error)
	UpdateUser(id uint64, input dto.UserCreate) error
	GetMembershipPackageByID(id uint64) (dto.MembershipPackageResponse, error)
}

type MembershipTransactionImplementation struct {
	repository membershiptransactionrepository.MembershipTransactionRepository
}

func (Mti *MembershipTransactionImplementation) GetAllMembershipTransaction() ([]dto.MembershipTransactionResponse, error) {
	data, err := Mti.repository.GetAllMembershipTransaction()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (Mti *MembershipTransactionImplementation) GetMembershipTransactionByID(id uint64) (dto.MembershipTransactionResponse, error) {
	data, err := Mti.repository.GetMembershipTransactionByID(id)
	if err != nil {
		return dto.MembershipTransactionResponse{}, err
	}

	var membershiptransaction dto.MembershipTransactionResponse
	err = copier.Copy(&membershiptransaction, &data)
	if err != nil {
		return dto.MembershipTransactionResponse{}, err
	}

	return membershiptransaction, nil
}

func (Mti *MembershipTransactionImplementation) CreateMembershipTransaction(input dto.MembershipTransactionCreate) error {
	err := Mti.repository.CreateMembershipTransaction(input)
	if err != nil {
		return err
	}
	return nil
}

func (Mti *MembershipTransactionImplementation) UpdateMembershipTransaction(id uint64, input dto.MembershipTransactionCreate) error {
	// call repository to update course
	err := Mti.repository.UpdateMembershipTransaction(id, input)
	if err != nil {
		return err
	}
	return nil
}

func (Mti *MembershipTransactionImplementation) DeleteMembershipTransaction(id uint64) error {
	err := Mti.repository.DeleteMembershipTransaction(id)
	if err != nil {
		return err
	}
	return nil
}

func (Mti *MembershipTransactionImplementation) GetUserByID(id uint64) (dto.UserCreate, error) {
	data, err := Mti.repository.GetUserByID(id)
	if err != nil {
		return dto.UserCreate{}, err
	}

	var user dto.UserCreate
	err = copier.Copy(&user, &data)
	if err != nil {
		return dto.UserCreate{}, err
	}

	return user, nil
}

func (Mti *MembershipTransactionImplementation) UpdateUser(id uint64, input dto.UserCreate) error {

	password := helpers.HashPassword(input.Password)

	input.Password = password

	// call repository to update course
	err := Mti.repository.UpdateUser(id, input)
	if err != nil {
		return err
	}

	return nil
}

func (Mti *MembershipTransactionImplementation) GetMembershipPackageByID(id uint64) (dto.MembershipPackageResponse, error) {
	data, err := Mti.repository.GetMembershipPackageByID(id)
	if err != nil {
		return dto.MembershipPackageResponse{}, err
	}

	var MembershipPackage dto.MembershipPackageResponse
	err = copier.Copy(&MembershipPackage, &data)
	if err != nil {
		return dto.MembershipPackageResponse{}, err
	}

	return MembershipPackage, nil
}

func NewMenuService(membershipTransactionRepo membershiptransactionrepository.MembershipTransactionRepository) MembershipTransactionService {
	return &MembershipTransactionImplementation{
		repository: membershipTransactionRepo,
	}
}
