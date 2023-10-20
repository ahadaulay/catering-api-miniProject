package dto

import (
	"time"

	"gorm.io/gorm"
)

type AdminResponse struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type AdminCreate struct{
	Name      string	`json:"name"`
	Email     string	`json:"email"`
	Password  string	`json:"password"`
}