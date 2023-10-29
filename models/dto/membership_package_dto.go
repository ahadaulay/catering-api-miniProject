package dto

import (
	"time"

	"gorm.io/gorm"
)

type MembershipPackageResponse struct {
	ID        uint64         `json:"id"`
	AdminID   uint64         `json:"admin_id"`
	Name      string         `json:"name"`
	Duration  uint32         `json:"duration"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type MembershipPackageCreate struct {
	ID       uint64 `json:"id"`
	AdminID  uint64 `json:"admin_id"`
	Name     string `json:"name"`
	Duration uint16 `json:"duration"`
}
