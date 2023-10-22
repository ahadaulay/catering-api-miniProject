package userservice

import (
	"catering-api/helpers"
	"catering-api/models/dto"
	userrepository "catering-api/repositorys/user_repository"

	"github.com/jinzhu/copier"
)

type UserService interface {
	GetAllUser() ([]dto.UserResponse, error)
	GetUserByID(id uint64) (dto.UserResponse, error)
	CreateUser(input dto.UserCreate) error
	UpdateUser(id uint64, input dto.UserCreate) error
	DeleteUser(id uint64) error
	LoginUser(input dto.UserLogin) (dto.UserResponse, error)
}

type UserImplementation struct{
	repository userrepository.UserRepository
}

func (Ui *UserImplementation) GetAllUser() ([]dto.UserResponse, error) {
	data, err := Ui.repository.GetAllUser()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (Ui *UserImplementation) GetUserByID(id uint64) (dto.UserResponse, error) {
	data, err := Ui.repository.GetUserByID(id)
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

func (Ui *UserImplementation) CreateUser(input dto.UserCreate) error {

	password, errPassword := helpers.HashPassword(input.Password)

	input.Password = password

	if errPassword != nil {
		return errPassword
	}

	err := Ui.repository.CreateUser(input)
	if err != nil {
		return err
	}
	return nil
}

func (Ui *UserImplementation) UpdateUser(id uint64, input dto.UserCreate) error {

	password, errPassword := helpers.HashPassword(input.Password)

	input.Password = password

	if errPassword != nil {
		return errPassword
	}

	// call repository to update course
	err := Ui.repository.UpdateUser(id,input)
	if err != nil {
		return err
	}

	return nil
}


func (Ui *UserImplementation) DeleteUser(id uint64) error {
	err := Ui.repository.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

func (Ui *UserImplementation) LoginUser(user dto.UserLogin) (dto.UserResponse, error) {
	// call repository to get user
	CostumerLogin, err := Ui.repository.LoginUser(user)
	if err != nil {
		return dto.UserResponse{}, err
	}
	return CostumerLogin, nil
}

func NewMenuService(userRepo userrepository.UserRepository) UserService {
	return &UserImplementation{
		repository: userRepo,
	}
}