package drivers

import (
	"catering-api/models/model"

	"gorm.io/gorm"
)

func DBMigration(db *gorm.DB) error {
	err := db.AutoMigrate(
		model.Admin{},
		model.User{},
		model.Menu{},
		model.Food{},
		model.Payment{},
		model.MembershipPackage{},
		model.MembershipTransaction{},
		model.MenuTransaction{},
	)

	if err != nil {
		return err
	}

	return nil
}