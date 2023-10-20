package model

import (
	"time"

	"gorm.io/gorm"
)

type MembershipPackage struct {
	ID 			uint64 `gorm:"primaryKey;notNull"`
	Name 		string
	Duration 	uint16
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	gorm.DeletedAt
	MembershipTransactions []MembershipTransaction
}