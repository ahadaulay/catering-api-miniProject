package drivers

import (
	"catering-api/models/model"

	"gorm.io/gorm"
)

func DBMigration(db *gorm.DB) error {
	err := db.AutoMigrate(
		model.Admin{},
		model.User{},
		model.Food{},
		model.MembershipPackage{},
		model.MembershipTransaction{},
		model.Menu{},
		model.MenuTransaction{},
		model.Payment{},
	)

	if err != nil {
		return err
	}

	return nil
}