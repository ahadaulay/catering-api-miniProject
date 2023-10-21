package adminrepository

import (
	"catering-api/models/dto"
	"catering-api/models/model"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type AdminImplementation struct {
	db *gorm.DB
}

func (Ai *AdminImplementation) GetAllAdmin() ([]dto.AdminResponse, error) {
	var AdminModel []model.Admin
	err := Ai.db.Find(&AdminModel).Error

	if err != nil {
		return nil, err
	}

	var Admin []dto.AdminResponse

	// Periksa jika ada error dalam proses penyalinan	
	if err := copier.Copy(&Admin, &AdminModel); err != nil {
		return nil, err
	}

	return Admin, nil
}


func (Ai *AdminImplementation) CreateAdmin(input dto.AdminCreate) (error)  {
	var AdminModel model.Admin

	err := copier.Copy(&AdminModel,&input)

	if err != nil {
		return err
	}

	err = Ai.db.Model(&model.Admin{}).Create(&AdminModel).Error

	if err != nil {
		return err
	}

	return nil
}

func NewAdminRepository(db *gorm.DB) AdminRepository  {
	return &AdminImplementation{
		db : db,
	}
}