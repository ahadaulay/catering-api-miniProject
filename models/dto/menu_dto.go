package dto

import (
	"time"

	"gorm.io/gorm"
)

type MenuResponse struct {
	ID        uint64 `json:"id"`
	AdminID   uint64 `json:"admin_id"`
	Name      string `json:"name"`
	MenuType  string `json:"menu_type"`
	Stock     int64  `json:"stock"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	DeletedAt gorm.DeletedAt	`json:"deleted_at"`
}

type MenuCreate struct {
	ID        uint64 `json:"id"`
	AdminID   uint64 `json:"admin_id"`
	Name      string `json:"name"`
	MenuType  string `json:"menu_type"`
	Stock     int64  `json:"stock"`
}
