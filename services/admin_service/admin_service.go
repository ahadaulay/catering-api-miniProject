package adminservice

import (
	"catering-api/helpers"
	"catering-api/models/dto"
	adminrepository "catering-api/repositorys/admin_repository"
)

type AdminService interface {
    GetAllAdmin() ([]dto.AdminResponse, error)
    CreateAdmin(input dto.AdminCreate) error
	LoginAdmin(admin dto.AdminLogin) (dto.AdminResponse,error)
}


type AdminImplementation struct{
	repository adminrepository.AdminRepository
}

func (Ai *AdminImplementation) GetAllAdmin() ([]dto.AdminResponse,error) {
	data ,err := Ai.repository.GetAllAdmin()

	if err != nil {
		return nil , err
	}

	return data , nil
}

func (Ai *AdminImplementation) CreateAdmin(input dto.AdminCreate) error  {
	hashpassword, err := helpers.HashPassword(input.Password)

	if err != nil {
		return err
	}

	input.Password = hashpassword

	err = Ai.repository.CreateAdmin(input)

	if err != nil {
		return err
	}

	return nil
}

func (Ai *AdminImplementation) LoginAdmin(admin dto.AdminLogin) (dto.AdminResponse, error) {
	// call repository to get user
	AdminLogin, err := Ai.repository.LoginAdmin(admin)
	if err != nil {
		return dto.AdminResponse{}, err
	}
	return AdminLogin, nil
}

func NewAdminService(AdminRepo adminrepository.AdminRepository) AdminService  {
    return &AdminImplementation{
        repository: AdminRepo,
    }
}



