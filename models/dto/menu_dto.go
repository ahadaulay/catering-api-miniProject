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
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type MenuCreate struct {
	ID        uint64 `json:"id"`
	AdminID   uint64 `json:"admin_id"`
	Name      string `json:"name"`
	MenuType  string `json:"menu_type"`
	Stock     int64  `json:"stock"`
}
