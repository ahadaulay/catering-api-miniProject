package dto

import (
	"time"

	"gorm.io/gorm"
)

type FoodResponse struct {
	ID        uint64		`json:"id"`
	AdminID   uint64		`json:"admin_id"`
	MenuID    uint64		`json:"menu_id"`
	Name      string		`json:"menu"`
	CreatedAt time.Time		
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type FoodCreate struct {
	ID        uint64		`json:"id"`
	AdminID   uint64		`json:"admin_id"`
	MenuID    uint64		`json:"menu_id"`
	Name      string		`json:"menu"`
}