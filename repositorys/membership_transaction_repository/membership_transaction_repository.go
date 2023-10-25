package membershiptransactionrepository

import "catering-api/models/dto"

type MembershipTransactionRepository interface {
	GetAllMembershipTransaction() ([]dto.MembershipTransactionResponse , error)
	GetMembershipTransactionByID(id uint64) (dto.MembershipTransactionResponse , error)
	CreateMembershipTransaction(input dto.MembershipTransactionCreate) error
	UpdateMembershipTransaction(id uint64 , input dto.MembershipTransactionCreate) error
	DeleteMembershipTransaction(id uint64) error
}