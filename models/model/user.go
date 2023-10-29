package model

import (
	"time"

	"gorm.io/gorm"
)



type User struct {
	ID                 uint64 `gorm:"primaryKey;notNull"`
	Name               string
	Email              string
	Password           string
	Address            string
	Phone              string
	Gender             string
	MembershipType     string
	MembershipDuration uint32
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt
	MenuTransactions	[]MenuTransaction
	MembershipTransactions []MembershipTransaction
}