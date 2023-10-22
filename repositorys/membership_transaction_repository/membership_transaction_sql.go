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

func (Mti *MembershipTransactionImplementation) GetMembershipTransactionByUserID(user_id uint64) (dto.MembershipTransactionResponse , error)  {
	var MembershipTransactionModel model.MembershipTransaction

	err := Mti.db.Model(&model.MembershipTransaction{}).Where("user_id = ? ", user_id).Find(&MembershipTransactionModel).Error

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





