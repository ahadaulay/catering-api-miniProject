package dto

import (
	"time"

	"gorm.io/gorm"
)

type PaymentResponse struct {
	ID            uint64		`json:"id"`
	AdminID       uint64		`json:"admin_id"`
	AccountName   string		`json:"account_name"`
	AccountNumber string		`json:"account_number"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	DeletedAt gorm.DeletedAt	`json:"deleted_at"`
}

type PaymentCreate struct {
	ID            uint64		`json:"id"`
	AdminID       uint64		`json:"admin_id"`
	AccountName   string		`json:"account_name"`
	AccountNumber string		`json:"account_number"`
}