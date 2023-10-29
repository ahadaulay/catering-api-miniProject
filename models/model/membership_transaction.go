package model

import (
	"time"

	"gorm.io/gorm"
)

type MembershipTransaction struct {
	ID                  uint64 `gorm:"primaryKey;notNull"`
	UserID              uint64
	MembershipPackageID uint64
	PaymentID           uint64
	Proof               string
	Status              string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt
}
