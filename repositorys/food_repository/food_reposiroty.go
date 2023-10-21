package foodrepository

import "catering-api/models/dto"

type FoodRepository interface {
	GetAllFood() ([]dto.FoodResponse,error)
	CreateFood(input dto.FoodCreate) error
	UpdateFood(input dto.FoodCreate) error
	DeleteFood(uint64) error
}