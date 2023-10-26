package dto

import (
	"time"

	"gorm.io/gorm"
)

type MembershipTransactionResponse struct {
	ID        			uint64			`json:"id"`
	MembershipPackageID uint64			`json:"membership_package_id"` 	
	UserID				uint64			`json:"user_id"`		
	PaymentID 			uint64			`json:"payment_id"`
	Status    			string			`json:"status"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	DeletedAt gorm.DeletedAt	`json:"deleted_at"`
}

type MembershipTransactionCreate struct {
	ID        			uint64			`json:"id"`
	UserID				uint64			`json:"user_id"`	
	MembershipPackageID uint64			`json:"membership_package_id"` 		
	PaymentID 			uint64			`json:"payment_id"`
	Status    			string			`json:"status"`
}