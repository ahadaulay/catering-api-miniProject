package foodservice

import (
	"catering-api/models/dto"
	foodrepository "catering-api/repositorys/food_repository"

	"github.com/jinzhu/copier"
)

type FoodService interface {
	GetAllFood() ([]dto.FoodResponse, error)
	GetFoodByID(uint64) (dto.FoodResponse, error)
	CreateFood(input dto.FoodCreate) error
	UpdateFood(id uint64 ,  input dto.FoodCreate) error
	DeleteFood(id uint64) error
}

type FoodImplementation struct {
	repository foodrepository.FoodRepository
}

func (Fi *FoodImplementation) GetAllFood() ([]dto.FoodResponse, error) {
	data, err := Fi.repository.GetAllFood()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (Fi *FoodImplementation) GetFoodByID(id uint64) (dto.FoodResponse, error) {
	data, err := Fi.repository.GetFoodByID(id)
	if err != nil {
		return dto.FoodResponse{}, err
	}

	var food dto.FoodResponse
	err = copier.Copy(&food, &data)
	if err != nil {
		return dto.FoodResponse{}, err
	}

	return food, nil
}

func (Fi *FoodImplementation) CreateFood(input dto.FoodCreate) error {
	err := Fi.repository.CreateFood(input)
	if err != nil {
		return err
	}
	return nil
}

func (Fi *FoodImplementation) UpdateFood(id uint64, input dto.FoodCreate) error {
	// call repository to update course
	err := Fi.repository.UpdateFood(id,input)
	if err != nil {
		return err
	}
	return nil
}


func (Fi *FoodImplementation) DeleteFood(id uint64) error {
	err := Fi.repository.DeleteFood(id)
	if err != nil {
		return err
	}
	return nil
}

func NewFoodService(foodRepo foodrepository.FoodRepository) FoodService {
	return &FoodImplementation{
		repository: foodRepo,
	}
}
