package model

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID uint64 `gorm:"primaryKey;notNull"`
	AccountName string
	AccountNumber string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}