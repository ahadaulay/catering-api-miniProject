package model

import (
	"catering-api/constants"
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
	Gender             constants.GenderType
	MembershipType     constants.MembershipType
	MembershipDuration uint32
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt
}