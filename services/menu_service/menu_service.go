package menuservice

import (
	"catering-api/models/dto"
	menurepository "catering-api/repositorys/menu_repository"

	"github.com/jinzhu/copier"
)

type MenuService interface {
	GetAllMenu() ([]dto.MenuResponse, error)
	GetMenuByID(uint64) (dto.MenuResponse, error)
	CreateMenu(input dto.MenuCreate) error
	UpdateMenu(id uint64, input dto.MenuResponse) error
	DeleteMenu(id uint64) error
	GetMenuByAdminID(adminID uint64) ([]dto.MenuResponse, error)
}

type MenuImplementation struct {
	repository menurepository.MenuRepository
}

func (Mi *MenuImplementation) GetAllMenu() ([]dto.MenuResponse, error) {
	data, err := Mi.repository.GetAllMenu()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (Mi *MenuImplementation) GetMenuByID(id uint64) (dto.MenuResponse, error) {
	data, err := Mi.repository.GetMenuByID(id)
	if err != nil {
		return dto.MenuResponse{}, err
	}

	var menu dto.MenuResponse
	err = copier.Copy(&menu, &data)
	if err != nil {
		return dto.MenuResponse{}, err
	}

	return menu, nil
}

func (Mi *MenuImplementation) CreateMenu(input dto.MenuCreate) error {
	err := Mi.repository.CreateMenu(input)
	if err != nil {
		return err
	}
	return nil
}

func (Mi *MenuImplementation) UpdateMenu(id uint64, input dto.MenuResponse) error {
	// call repository to update course
	err := Mi.repository.UpdateMenu(id, input)
	if err != nil {
		return err
	}
	return nil
}

func (Mi *MenuImplementation) DeleteMenu(id uint64) error {
	err := Mi.repository.DeleteMenu(id)
	if err != nil {
		return err
	}
	return nil
}

func (Mi *MenuImplementation) GetMenuByAdminID(adminID uint64) ([]dto.MenuResponse, error) {
	data, err := Mi.repository.GetMenuByAdminID(adminID)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewMenuService(menuRepo menurepository.MenuRepository) MenuService {
	return &MenuImplementation{
		repository: menuRepo,
	}
}
