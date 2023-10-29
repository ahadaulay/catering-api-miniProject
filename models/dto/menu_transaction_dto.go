package dto

import (
	"gorm.io/gorm"
	"time"
)

type MenuTransactionResponse struct {
	ID        uint64         `json:"id"`
	UserID    uint64         `json:"user_id"`
	MenuID    uint64         `json:"menu_id"`
	Status    string         `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type MenuTransactionCreate struct {
	UserID uint64 `json:"user_id"`
	MenuID uint64 `json:"menu_id"`
	Status string `json:"status"`
}
