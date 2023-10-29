package adminservice

import (
	"catering-api/helpers"
	req "catering-api/helpers/request"
	"catering-api/models/dto"
	"catering-api/models/model"
	adminrepository "catering-api/repositorys/admin_repository"
	"fmt"
	"github.com/labstack/echo/v4"
)

type AdminService interface {
	GetAllAdmin() ([]dto.AdminResponse, error)
	CreateAdmin(input dto.AdminCreate) error
	LoginAdmin(c echo.Context, request dto.AdminLogin) (*model.Admin, error)
}

type AdminImplementation struct {
	repository adminrepository.AdminRepository
}

func (Ai *AdminImplementation) GetAllAdmin() ([]dto.AdminResponse, error) {
	data, err := Ai.repository.GetAllAdmin()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (Ai *AdminImplementation) CreateAdmin(input dto.AdminCreate) error {
	hashpassword := helpers.HashPassword(input.Password)

	input.Password = hashpassword

	err := Ai.repository.CreateAdmin(input)

	if err != nil {
		return err
	}

	return nil
}

func (Ai *AdminImplementation) LoginAdmin(c echo.Context, request dto.AdminLogin) (*model.Admin, error) {

	existingDoctor, err := Ai.repository.FindByEmail(request.Email)

	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	doctor := req.AdminLoginRequestToAdminDomain(request)

	err = helpers.ComparePassword(existingDoctor.Password, doctor.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	return existingDoctor, nil

}

func NewAdminService(AdminRepo adminrepository.AdminRepository) AdminService {
	return &AdminImplementation{
		repository: AdminRepo,
	}
}
