package dto

import (
	"time"

	"gorm.io/gorm"
)

type MembershipTransactionResponse struct {
	ID        uint64			`json:"id"`
	UserID    uint64			`json:"user_id"`
	MenuID    uint64			`json:"menu_id"`
	Status    string			`json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type MembershipTransactionCreate struct {
	ID        uint64			`json:"id"`
	UserID    uint64			`json:"user_id"`
	MenuID    uint64			`json:"menu_id"`
	Status    string			`json:"status"`
}