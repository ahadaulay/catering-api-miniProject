package userrepository

import "catering-api/models/dto"

type UserRepository interface {
	GetAllUser() ([]dto.UserResponse,error)
	GetUserByID(id uint64)(dto.UserResponse,error)
	CreateUser(input dto.UserCreate) error
	UpdateUser(id uint64,input dto.UserCreate) error
	DeleteUser(id uint64) error
	LoginUser(input dto.UserLogin) (dto.UserResponse , error)
}