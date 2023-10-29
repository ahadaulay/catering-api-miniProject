package adminrepository

import (
	"catering-api/helpers"
	"catering-api/models/dto"
	"catering-api/models/model"
	"errors"

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

func (Ai *AdminImplementation) CreateAdmin(input dto.AdminCreate) error {
	var AdminModel model.Admin

	err := copier.Copy(&AdminModel, &input)

	if err != nil {
		return err
	}

	err = Ai.db.Model(&model.Admin{}).Create(&AdminModel).Error

	if err != nil {
		return err
	}

	return nil
}

func (Ai *AdminImplementation) LoginAdmin(input dto.AdminLogin) (dto.AdminResponse, error) {
	var AdminLogin dto.AdminResponse

	err := Ai.db.Model(&model.Admin{}).First(&AdminLogin, "email = ?", input.Email).Error

	if err != nil {
		return dto.AdminResponse{}, errors.New("email not registered")
	}

	err = helpers.ComparePassword(input.Password, AdminLogin.Password)

	if err != nil {
		return dto.AdminResponse{}, errors.New("wrong password")
	}
	var AdminLoginResponse = dto.AdminResponse{
		ID:       AdminLogin.ID,
		Name:     AdminLogin.Name,
		Email:    AdminLogin.Email,
		Password: AdminLogin.Password,
	}

	return AdminLoginResponse, nil
}

func (Ai *AdminImplementation) FindByEmail(email string) (*model.Admin, error) {
	admin := model.Admin{}

	result := Ai.db.Where("email = ?", email).First(&admin)
	if result.Error != nil {
		return nil, result.Error
	}

	return &admin, nil
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &AdminImplementation{
		db: db,
	}
}
