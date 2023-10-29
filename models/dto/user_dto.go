package dto

import (
	"time"

	"gorm.io/gorm"
)

type UserResponse struct {
	ID                 uint64		`json:"id"`
	Name               string		`json:"name"`
	Email              string		`json:"email"`
	Password           string		`json:"password"`
	Address            string		`json:"address"`
	Phone              string		`json:"phone"`
	Gender             string		`json:"gender"`
	MembershipType     string		`json:"membership_type"`
	MembershipDuration uint32		`json:"membership_duration"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	DeletedAt gorm.DeletedAt	`json:"deleted_at"`
}

type UserCreate struct {
	ID                 uint64		`json:"id"`
	Name               string		`json:"name"`
	Email              string		`json:"email"`
	Password           string		`json:"password"`
	Address            string		`json:"address"`
	Phone              string		`json:"phone"`
	Gender             string		`json:"gender"`
	MembershipType     string		`json:"membership_type"`
	MembershipDuration uint32		`json:"membership_duration"`
}

type UserLogin struct{
	ID                 uint64		`json:"id"`
	Email              string		`json:"email"`
	Password           string		`json:"password"`
}

type UserResponseLogin struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}