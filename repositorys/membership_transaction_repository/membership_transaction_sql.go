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

func (Mti *MembershipTransactionImplementation) GetAllMembershipTransaction() ([]dto.MembershipTransactionResponse,error)  {
	var MembershipTransactionModel []model.MembershipTransaction

	err := Mti.db.Find(&MembershipTransactionModel).Error

	if err != nil {
		return nil, err
	}

	var MembershipTransaction []dto.MembershipTransactionResponse

	if err := copier.Copy(&MembershipTransaction, &MembershipTransactionModel); err != nil {
		return nil, err
	}

	return MembershipTransaction , nil
}

func (Mti *MembershipTransactionImplementation) GetMembershipTransactionByID(id uint64) (dto.MembershipTransactionResponse , error)  {
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

	err = copier.Copy(&MembershipTransaction,&MembershipTransactionModel)

	if err != nil {
		return dto.MembershipTransactionResponse{}, err
	}

	return MembershipTransaction, nil 
}

func (Mti *MembershipTransactionImplementation) CreateMembershipTransaction(input dto.MembershipTransactionCreate) error  {
	var MembershipTransactionModel model.MembershipTransaction

	MembershipTransactionModel.Status = "pending"

	err := copier.Copy(&MembershipTransactionModel,&input)

	if err != nil {
		return err
	}

	err = Mti.db.Model(&model.MembershipTransaction{}).Create(&MembershipTransactionModel).Error

	if err != nil {
		return err
	}

	return nil
}

func (Mti *MembershipTransactionImplementation) UpdateMembershipTransaction(id uint64, input dto.MembershipTransactionCreate) error {
	// Update menu with new data
	result := Mti.db.Model(&model.MembershipTransaction{}).Where("id = ?", id).Updates(&model.MembershipTransaction{
		MembershipPackageID: input.ID ,
		PaymentID: input.PaymentID,
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


func (Mti *MembershipTransactionImplementation) DeleteMembershipTransaction(id uint64) error {
	err := Mti.db.Where("id = ?", id).Delete(&model.MembershipTransaction{}).Error

	if err != nil {
		return err
	}

	return nil
}

func NewMenuRepository(db *gorm.DB) MembershipTransactionRepository  {
	return &MembershipTransactionImplementation{
		db: db,
	}
}




