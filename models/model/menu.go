package model

import (
	"catering-api/constants"
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	ID       	uint64
	AdminID  	uint64
	Name     	string
	MenuType 	constants.MembershipType
	Stock 		int64
	CreatedAt	time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}