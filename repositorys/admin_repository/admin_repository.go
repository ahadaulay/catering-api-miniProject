package adminrepository

import (
	"catering-api/models/dto"
	"catering-api/models/model"
)

type AdminRepository interface {
	GetAllAdmin() ([]dto.AdminResponse, error)
	CreateAdmin(input dto.AdminCreate) error
	LoginAdmin(input dto.AdminLogin) (dto.AdminResponse, error)
	FindByEmail(email string) (*model.Admin, error)
}
