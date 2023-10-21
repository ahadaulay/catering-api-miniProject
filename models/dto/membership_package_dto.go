package dto

import (
	"time"

	"gorm.io/gorm"
)

type MembershipPackageResponse struct {
	ID                  uint64 	`json:"id"`
	AdminID				uint64	`json:"admin_id"`
	Name                string	`json:"name"`
	Duration            uint16	`json:"duration"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt
}

type MembershipPackageCreate struct {
	ID                  uint64 	`json:"id"`
	AdminID				uint64	`json:"admin_id"`
	Name                string	`json:"name"`
	Duration            uint16	`json:"duration"`
}