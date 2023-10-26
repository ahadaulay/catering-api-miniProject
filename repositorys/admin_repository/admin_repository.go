package adminrepository

import "catering-api/models/dto"

type AdminRepository interface {
	GetAllAdmin() ([]dto.AdminResponse,error)
	CreateAdmin(input dto.AdminCreate) (error)
	LoginAdmin(input dto.AdminLogin) (dto.AdminResponse , error)
}