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
	Image	  string		`json:"image"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	DeletedAt gorm.DeletedAt	`json:"deleted_at"`
}

type FoodCreate struct {
	ID        uint64		`json:"id"`
	AdminID   uint64		`json:"admin_id"`
	MenuID    uint64		`json:"menu_id"`
	Name      string		`json:"menu"`
	Image	  string		`json:"image"`
}