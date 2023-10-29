package menurepository

import (
	"catering-api/models/dto"
	"catering-api/models/model"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type MenuImplementation struct {
	db *gorm.DB
}

func (Mi *MenuImplementation) GetAllMenu() ([]dto.MenuResponse, error) {
	var MenuModel []model.Menu

	err := Mi.db.Find(&MenuModel).Error

	if err != nil {
		return nil, err
	}

	var Menu []dto.MenuResponse

	if err := copier.Copy(&Menu, &MenuModel); err != nil {
		return nil, err
	}

	return Menu, nil
}

func (Mi *MenuImplementation) GetMenuByID(id uint64) (dto.MenuResponse, error) {
	var MenuModel model.Menu

	err := Mi.db.Model(&model.Menu{}).Where("id = ? ", id).Find(&MenuModel).Error

	if err != nil {
		return dto.MenuResponse{}, err
	}

	// Periksa apakah data ditemukan
	if MenuModel.ID == 0 {
		return dto.MenuResponse{}, gorm.ErrRecordNotFound
	}

	var Menu dto.MenuResponse

	err = copier.Copy(&Menu, &MenuModel)

	if err != nil {
		return dto.MenuResponse{}, err
	}

	return Menu, nil
}

func (Mi *MenuImplementation) CreateMenu(input dto.MenuCreate) error {
	var MenuModel model.Menu

	err := copier.Copy(&MenuModel, &input)

	if err != nil {
		return err
	}

	err = Mi.db.Model(&model.Menu{}).Create(&MenuModel).Error

	if err != nil {
		return err
	}

	return nil
}

func (Mi *MenuImplementation) UpdateMenu(id uint64, input dto.MenuResponse) error {
	// Update menu with new data
	result := Mi.db.Model(&model.Menu{}).Where("id = ?", id).Updates(&model.Menu{
		Name:     input.Name,
		AdminID:  input.AdminID,
		MenuType: input.MenuType,
		Stock:    input.Stock,
	})

	if result.Error != nil {
		return result.Error
	}

	// Check if any rows were affected
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (Mi *MenuImplementation) DeleteMenu(id uint64) error {
	err := Mi.db.Where("id = ?", id).Delete(&model.Menu{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (Mi *MenuImplementation) GetMenuByAdminID(adminID uint64) ([]dto.MenuResponse, error) {
	var MenuModel []model.Menu

	err := Mi.db.Model(&model.Menu{}).Where("admin_id = ? ", adminID).Find(&MenuModel).Error

	// Periksa apakah FoodModel kosong (tidak ada data yang ditemukan)
	if len(MenuModel) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	var Menu []dto.MenuResponse

	err = copier.Copy(&Menu, &MenuModel)

	if err != nil {
		return nil, err
	}

	return Menu, nil
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &MenuImplementation{
		db: db,
	}
}
