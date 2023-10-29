package userservice

import (
	"catering-api/helpers"
	req "catering-api/helpers/request"
	"catering-api/models/dto"
	"catering-api/models/model"
	userrepository "catering-api/repositorys/user_repository"
	"fmt"
	"github.com/labstack/echo/v4"

	"github.com/jinzhu/copier"
)

type UserService interface {
	GetAllUser() ([]dto.UserResponse, error)
	GetUserByID(id uint64) (dto.UserResponse, error)
	CreateUser(input dto.UserCreate) error
	UpdateUser(id uint64, input dto.UserCreate) error
	DeleteUser(id uint64) error
	LoginUser(ctx echo.Context, request dto.UserLogin) (*model.User, error)
}

type UserImplementation struct {
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

	password := helpers.HashPassword(input.Password)

	input.Password = password
	input.MembershipDuration = 0

	err := Ui.repository.CreateUser(input)
	if err != nil {
		return err
	}
	return nil
}

func (Ui *UserImplementation) UpdateUser(id uint64, input dto.UserCreate) error {

	password := helpers.HashPassword(input.Password)

	input.Password = password

	// call repository to update course
	err := Ui.repository.UpdateUser(id, input)
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

func (Ui *UserImplementation) LoginUser(c echo.Context, request dto.UserLogin) (*model.User, error) {

	existingUser, err := Ui.repository.FindByEmail(request.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	user := req.UserLoginRequestToUserDomain(request)

	err = helpers.ComparePassword(existingUser.Password, user.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	return existingUser, nil

}

func NewMenuService(userRepo userrepository.UserRepository) UserService {
	return &UserImplementation{
		repository: userRepo,
	}
}
