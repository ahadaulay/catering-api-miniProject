package menutransactionrepository

import (
	"catering-api/models/dto"
	"catering-api/models/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type MenuTransactionImplementation struct {
	db *gorm.DB
}

func (Mti *MenuTransactionImplementation) GetAllMenuTransaction() ([]dto.MenuTransactionResponse, error) {
	var MenuTransactionModel []model.MenuTransaction

	err := Mti.db.Find(&MenuTransactionModel).Error

	if err != nil {
		return nil, err
	}

	var MenuTransaction []dto.MenuTransactionResponse

	if err := copier.Copy(&MenuTransaction, &MenuTransactionModel); err != nil {
		return nil, err
	}

	return MenuTransaction, nil
}

func (Mti *MenuTransactionImplementation) GetMenuTransactionByID(id uint64) (dto.MenuTransactionResponse, error) {
	var MenuTransactionModel model.MenuTransaction

	err := Mti.db.Model(&model.MenuTransaction{}).Where("id = ? ", id).Find(&MenuTransactionModel).Error

	if err != nil {
		return dto.MenuTransactionResponse{}, err
	}

	// Periksa apakah data ditemukan
	if MenuTransactionModel.ID == 0 {
		return dto.MenuTransactionResponse{}, gorm.ErrRecordNotFound
	}

	var MenuTransaction dto.MenuTransactionResponse

	err = copier.Copy(&MenuTransaction, &MenuTransactionModel)

	if err != nil {
		return dto.MenuTransactionResponse{}, err
	}

	return MenuTransaction, nil
}

func (Mti *MenuTransactionImplementation) CreateMenuTransaction(input dto.MenuTransactionCreate) error {
	var MenuTransactionModel model.MenuTransaction

	err := copier.Copy(&MenuTransactionModel, &input)

	if err != nil {
		return err
	}

	MenuTransactionModel.Status = "pending"

	err = Mti.db.Model(&model.MenuTransaction{}).Create(&MenuTransactionModel).Error

	if err != nil {
		return err
	}

	return nil
}

func (Mti *MenuTransactionImplementation) UpdateMenuTransaction(id uint64, input dto.MenuTransactionCreate) error {
	// Update menu with new data
	result := Mti.db.Model(&model.MenuTransaction{}).Where("id = ?", id).Updates(&model.MenuTransaction{
		UserID: input.UserID,
		MenuID: input.MenuID,
		Status: input.Status,
	})

	if result.Error != nil {
		return result.Error
	}

	// Check if any rows were affected
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (Mti *MenuTransactionImplementation) DeleteMenuTransaction(id uint64) error {
	err := Mti.db.Where("id = ?", id).Delete(&model.MenuTransaction{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (Mti *MenuTransactionImplementation) GetMenuByID(id uint64) (dto.MenuResponse, error) {
	var MenuModel model.Menu

	err := Mti.db.Model(&model.Menu{}).Where("id = ? ", id).Find(&MenuModel).Error

	if err != nil {
		return dto.MenuResponse{}, err
	}

	// Periksa apakah data ditemukan
	if MenuModel.ID == 0 {
		return dto.MenuResponse{}, gorm.ErrRecordNotFound
	}

	var Menu dto.MenuResponse

	err = copier.Copy(&Menu, &MenuModel)

	if err != nil {
		return dto.MenuResponse{}, err
	}

	return Menu, nil
}

func (Mti *MenuTransactionImplementation) GetUserByID(id uint64) (dto.UserResponse, error) {
	var UserModel model.User

	err := Mti.db.Model(&model.User{}).Where("id = ? ", id).Find(&UserModel).Error

	if err != nil {
		return dto.UserResponse{}, err
	}

	// Periksa apakah data ditemukan
	if UserModel.ID == 0 {
		return dto.UserResponse{}, gorm.ErrRecordNotFound
	}

	var User dto.UserResponse

	err = copier.Copy(&User, &UserModel)

	if err != nil {
		return dto.UserResponse{}, err
	}

	return User, nil
}

func (Mti *MenuTransactionImplementation) ReduceMenuStock(id uint64) error {
	// Make sure the database connection is properly initialized
	var input model.Menu

	err := Mti.db.Model(&model.Menu{}).Where("id = ? ", id).Find(&input).Error

	if err != nil {
		return err
	}

	// Periksa apakah data ditemukan
	if input.ID == 0 {
		return gorm.ErrRecordNotFound
	}

	// Start a transaction
	tx := Mti.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Fetch the current menu
	var menu model.Menu
	if err := tx.First(&menu, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Update the menu with new data
	menu.Name = input.Name
	menu.AdminID = input.AdminID
	menu.MenuType = input.MenuType

	// Ensure stock doesn't go below 0
	if input.Stock > 0 {
		menu.Stock = input.Stock - 1
	} else {
		menu.Stock = 0
	}

	if err := tx.Save(&menu).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	tx.Commit()
	return nil
}

func (Mti *MenuTransactionImplementation) ReduceDurationUser(id uint64) error {
	// Make sure the database connection is properly initialized
	var input model.User

	err := Mti.db.Model(&model.User{}).Where("id = ? ", id).Find(&input).Error

	if err != nil {
		return err
	}

	// Periksa apakah data ditemukan
	if input.ID == 0 {
		return gorm.ErrRecordNotFound
	}

	// Start a transaction
	tx := Mti.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Fetch the current menu
	var user model.User
	if err := tx.First(&user, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Ensure stock doesn't go below 0
	if input.MembershipDuration > 0 {
		user.MembershipDuration = input.MembershipDuration - 1
	} else {
		user.MembershipDuration = 0
	}

	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	tx.Commit()
	return nil
}

func NewMenuTransactionRepository(db *gorm.DB) MenuTransactionRepository {
	return &MenuTransactionImplementation{
		db: db,
	}
}
