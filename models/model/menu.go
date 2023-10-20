package model

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	ID       			uint64
	AdminID  			uint64
	Name     			string
	MenuType 			string
	Stock 				int64
	CreatedAt			time.Time
	UpdatedAt   		time.Time
	DeletedAt   		gorm.DeletedAt
	MenuTransactions 	[]MenuTransaction
	Foods				[]Food
}