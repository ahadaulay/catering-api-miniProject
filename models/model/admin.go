package model

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        uint64 `gorm:"primaryKey;notNull"`
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}