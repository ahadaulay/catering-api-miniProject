package dto

import (
	"time"

	"gorm.io/gorm"
)

type MembershipTransactionResponse struct {
	ID        			uint64			`json:"id"`
	MembershipPackageID uint64			`json:"membership_package_id"` 		
	PaymentID 			uint64			`json:"payment_id"`
	Status    			string			`json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type MembershipTransactionCreate struct {
	ID        			uint64			`json:"id"`
	MembershipPackageID uint64			`json:"membership_package_id"` 		
	PaymentID 			uint64			`json:"payment_id"`
	Status    			string			`json:"status"`
}