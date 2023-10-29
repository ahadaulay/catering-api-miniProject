package membershiptransactionrepository

import (
	"catering-api/models/dto"
	"catering-api/models/model"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type MembershipTransactionImplementation struct {
	db *gorm.DB
}

func (Mti *MembershipTransactionImplementation) GetAllMembershipTransaction() ([]dto.MembershipTransactionResponse, error) {
	var MembershipTransactionModel []model.MembershipTransaction

	err := Mti.db.Find(&MembershipTransactionModel).Error

	if err != nil {
		return nil, err
	}

	var MembershipTransaction []dto.MembershipTransactionResponse

	if err := copier.Copy(&MembershipTransaction, &MembershipTransactionModel); err != nil {
		return nil, err
	}

	return MembershipTransaction, nil
}

func (Mti *MembershipTransactionImplementation) GetMembershipTransactionByID(id uint64) (dto.MembershipTransactionResponse, error) {
	var MembershipTransactionModel model.MembershipTransaction

	err := Mti.db.Model(&model.MembershipTransaction{}).Where("id = ? ", id).Find(&MembershipTransactionModel).Error

	if err != nil {
		return dto.MembershipTransactionResponse{}, err
	}

	// Periksa apakah data ditemukan
	if MembershipTransactionModel.ID == 0 {
		return dto.MembershipTransactionResponse{}, gorm.ErrRecordNotFound
	}

	var MembershipTransaction dto.MembershipTransactionResponse

	err = copier.Copy(&MembershipTransaction, &MembershipTransactionModel)

	if err != nil {
		return dto.MembershipTransactionResponse{}, err
	}

	return MembershipTransaction, nil
}

func (Mti *MembershipTransactionImplementation) CreateMembershipTransaction(input dto.MembershipTransactionCreate) error {
	var MembershipTransactionModel model.MembershipTransaction

	err := copier.Copy(&MembershipTransactionModel, &input)

	if err != nil {
		return err
	}

	MembershipTransactionModel.Status = "pending"

	err = Mti.db.Model(&model.MembershipTransaction{}).Create(&MembershipTransactionModel).Error

	if err != nil {
		return err
	}

	return nil
}

func (Mti *MembershipTransactionImplementation) UpdateMembershipTransaction(id uint64, input dto.MembershipTransactionCreate) error {
	// Update menu with new data
	result := Mti.db.Model(&model.MembershipTransaction{}).Where("id = ?", id).Updates(&model.MembershipTransaction{
		MembershipPackageID: input.MembershipPackageID,
		UserID:              input.UserID,
		PaymentID:           input.PaymentID,
		Status:              input.Status,
		Proof:               input.Proof,
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

func (Mti *MembershipTransactionImplementation) DeleteMembershipTransaction(id uint64) error {
	err := Mti.db.Where("id = ?", id).Delete(&model.MembershipTransaction{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (Mti *MembershipTransactionImplementation) UpdateUser(id uint64, input dto.UserCreate) error {
	// Update menu with new data
	result := Mti.db.Model(&model.User{}).Where("id = ?", id).Updates(&model.User{
		Name:               input.Name,
		Email:              input.Email,
		Password:           input.Password,
		Address:            input.Address,
		Phone:              input.Phone,
		Gender:             input.Gender,
		MembershipType:     input.MembershipType,
		MembershipDuration: input.MembershipDuration,
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

func (Mti *MembershipTransactionImplementation) GetUserByID(id uint64) (dto.UserCreate, error) {
	var UserModel model.User

	err := Mti.db.Model(&model.User{}).Where("id = ? ", id).Find(&UserModel).Error

	if err != nil {
		return dto.UserCreate{}, err
	}

	// Periksa apakah data ditemukan
	if UserModel.ID == 0 {
		return dto.UserCreate{}, gorm.ErrRecordNotFound
	}

	var User dto.UserCreate

	err = copier.Copy(&User, &UserModel)

	if err != nil {
		return dto.UserCreate{}, err
	}

	return User, nil
}

func (Mti *MembershipTransactionImplementation) GetMembershipPackageByID(id uint64) (dto.MembershipPackageResponse, error) {
	var MembershipPackageModel model.MembershipPackage

	err := Mti.db.Model(&model.MembershipPackage{}).Where("id = ? ", id).Find(&MembershipPackageModel).Error

	if err != nil {
		return dto.MembershipPackageResponse{}, err
	}

	// Periksa apakah data ditemukan
	if MembershipPackageModel.ID == 0 {
		return dto.MembershipPackageResponse{}, gorm.ErrRecordNotFound
	}

	var MembershipPackage dto.MembershipPackageResponse

	err = copier.Copy(&MembershipPackage, &MembershipPackageModel)

	if err != nil {
		return dto.MembershipPackageResponse{}, err
	}

	return MembershipPackage, nil
}

func NewMenuRepository(db *gorm.DB) MembershipTransactionRepository {
	return &MembershipTransactionImplementation{
		db: db,
	}
}
