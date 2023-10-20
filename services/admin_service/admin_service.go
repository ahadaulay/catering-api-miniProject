package adminservice

import (
	"catering-api/helpers"
	"catering-api/models/dto"
	adminrepository "catering-api/repositorys/admin_repository"
)

type AdminService interface {
    GetAllAdmin() ([]dto.AdminResponse, error)
    CreateAdmin(input dto.AdminCreate) error
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

func NewAdminService(AdminRepo adminrepository.AdminRepository) AdminService  {
    return &AdminImplementation{
        repository: AdminRepo,
    }
}



