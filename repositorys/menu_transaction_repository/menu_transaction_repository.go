package menutransactionrepository

import "catering-api/models/dto"

type MenuTransactionRepository interface {
	GetAllMenuTransaction() ([]dto.MenuTransactionResponse, error)
	GetMenuTransactionByID(id uint64) (dto.MenuTransactionResponse, error)
	CreateMenuTransaction(input dto.MenuTransactionCreate) error
	UpdateMenuTransaction(id uint64, input dto.MenuTransactionCreate) error
	DeleteMenuTransaction(id uint64) error
	ReduceMenuStock(id uint64) error
	GetMenuByID(id uint64) (dto.MenuResponse, error)
	GetUserByID(id uint64) (dto.UserResponse, error)
	ReduceDurationUser(id uint64) error
}
