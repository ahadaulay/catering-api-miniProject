package foodrepository

import (
	"catering-api/models/dto"
	"catering-api/models/model"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type FoodImplementation struct {
	db *gorm.DB
}

func (Fi *FoodImplementation) GetAllFood() ([]dto.FoodResponse, error) {
	var FoodModel []model.Food

	err := Fi.db.Find(&FoodModel).Error

	if err != nil {
		return nil, err
	}

	var Food []dto.FoodResponse

	if err := copier.Copy(&Food, &FoodModel); err != nil {
		return nil, err
	}

	return Food, nil
}

func (Fi *FoodImplementation) GetFoodByID(id uint64) (dto.FoodResponse, error) {
	var FoodModel model.Food

	err := Fi.db.Model(&model.Food{}).Where("id = ? ", id).Find(&FoodModel).Error

	if err != nil {
		return dto.FoodResponse{}, err
	}

	// Periksa apakah data ditemukan
	if FoodModel.ID == 0 {
		return dto.FoodResponse{}, gorm.ErrRecordNotFound
	}

	var Food dto.FoodResponse

	err = copier.Copy(&Food, &FoodModel)

	if err != nil {
		return dto.FoodResponse{}, err
	}

	return Food, nil
}

func (Fi *FoodImplementation) CreateFood(input dto.FoodCreate) error {
	var FoodModel model.Food

	err := copier.Copy(&FoodModel, &input)

	if err != nil {
		return err
	}

	err = Fi.db.Model(&model.Food{}).Create(&FoodModel).Error

	if err != nil {
		return err
	}

	return nil
}

func (Fi *FoodImplementation) UpdateFood(id uint64, input dto.FoodCreate) error {
	var FoodModel model.Food

	err := copier.Copy(&FoodModel, &input)

	if err != nil {
		return err
	}

	err = Fi.db.Model(&model.Food{}).Where("id=?", id).Updates(&FoodModel).Error

	if err != nil {
		return err
	}

	return nil

}

func (Fi *FoodImplementation) DeleteFood(id uint64) error {
	err := Fi.db.Where("id = ?", id).Delete(&model.Food{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (Fi *FoodImplementation) GetFoodByAdminID(adminID uint64) ([]dto.FoodResponse, error) {
	var FoodModel []model.Food

	err := Fi.db.Model(&model.Food{}).Where("admin_id = ? ", adminID).Find(&FoodModel).Error

	// Periksa apakah FoodModel kosong (tidak ada data yang ditemukan)
	if len(FoodModel) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	var Food []dto.FoodResponse

	err = copier.Copy(&Food, &FoodModel)

	if err != nil {
		return nil, err
	}

	return Food, nil
}

func NewFoodRepository(db *gorm.DB) FoodRepository {
	return &FoodImplementation{
		db: db,
	}
}
