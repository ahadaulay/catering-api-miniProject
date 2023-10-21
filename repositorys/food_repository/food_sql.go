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

func (Fi *FoodImplementation) GetAllFood() ([]dto.FoodResponse , error)  {
	var FoodModel model.Food

	err := Fi.db.Find(&FoodModel).Error

	if err != nil {
		return nil, err
	}

	var Food []dto.FoodResponse

	if err := copier.Copy(&Food, &FoodModel); err != nil {
		return nil, err
	}

	return Food , nil
}

func (Fi *FoodImplementation) CreateFood(input dto.FoodCreate) (error)  {
	var FoodModel model.Food

	err := copier.Copy(&FoodModel,&input)

	if err != nil {
		return err
	}

	err = Fi.db.Model(&model.Food{}).Create(&FoodModel).Error

	if err != nil {
		return err
	}

	return nil
}

func (Fi *FoodImplementation) UpdateFood(input dto.FoodCreate) error  {
	var FoodModel model.Food

	err := copier.Copy(&FoodModel,&input)

	if err != nil {
		return err
	}

	err = Fi.db.Model(&model.Food{}).Where("id=?", FoodModel.ID).Updates(&FoodModel).Error

	if err!= nil {
		return err
	}

	return nil

}

func (Fi *FoodImplementation) DeleteFood(id int64) error {
	err := Fi.db.Where("id = ?", id).Delete(&model.Food{}).Error

	if err != nil {
		return err
	}

	return nil
}



