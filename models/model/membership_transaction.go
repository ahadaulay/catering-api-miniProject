package model

import (
	"catering-api/constants"
	"time"

	"gorm.io/gorm"
)

type MembershipTransaction struct {
	ID                  uint64 `gorm:"primaryKey;notNull"`
	MembershipPackageID uint64
	PaymentID 			uint64
	Status				constants.StatusTransactions
	CreatedAt           time.Time
	UpdatedAt           time.Time	
	DeletedAt           gorm.DeletedAt
}