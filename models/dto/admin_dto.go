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
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	DeletedAt gorm.DeletedAt	`json:"deleted_at"`
}

type AdminCreate struct{
	Name      string	`json:"name"`
	Email     string	`json:"email"`
	Password  string	`json:"password"`
}