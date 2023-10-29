package foodrepository

import "catering-api/models/dto"

type FoodRepository interface {
	GetAllFood() ([]dto.FoodResponse, error)
	GetFoodByID(id uint64) (dto.FoodResponse, error)
	CreateFood(input dto.FoodCreate) error
	UpdateFood(id uint64, input dto.FoodCreate) error
	DeleteFood(uint64) error
	GetFoodByAdminID(adminID uint64) ([]dto.FoodResponse, error)
}
