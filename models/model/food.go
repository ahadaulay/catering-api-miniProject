package model

import (
	"time"

	"gorm.io/gorm"
)

type Food struct {
	ID        	uint64 `gorm:"primaryKey;notNull"`
	AdminID 	uint64
	MenuID    	uint64
	Name     	string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	gorm.DeletedAt
}