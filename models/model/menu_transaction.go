package model

import (
	"catering-api/constants"
	"time"

	"gorm.io/gorm"
)

type MenuTransaction struct {
	ID     		uint64 `gorm:"primaryKey;notNull"`
	UserID 		uint64
	MenuID 		uint64
	Status 		constants.StatusTransactions
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt

}