package menurepository

import "catering-api/models/dto"

type MenuRepository interface {
	GetAllMenu() ([]dto.MenuResponse, error)
	GetMenuByID(uint64) (dto.MenuResponse, error)
	CreateMenu(input dto.MenuCreate) error
	UpdateMenu(id uint64, input dto.MenuResponse) error
	DeleteMenu(id uint64) error
	GetMenuByAdminID(adminID uint64) ([]dto.MenuResponse, error)
}
