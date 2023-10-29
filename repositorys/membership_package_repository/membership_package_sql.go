package membershippackagerepository

import (
	"catering-api/models/dto"
	"catering-api/models/model"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type MembershipPackageImplementation struct {
	db *gorm.DB
}

func (Mpi MembershipPackageImplementation) GetAllMembershipPackage() ([]dto.MembershipPackageResponse, error) {
	var MembershipPackageModel []model.MembershipPackage

	err := Mpi.db.Find(&MembershipPackageModel).Error

	if err != nil {
		return nil, err
	}

	var MembershipPackage []dto.MembershipPackageResponse

	if err := copier.Copy(&MembershipPackage, &MembershipPackageModel); err != nil {
		return nil, err
	}

	return MembershipPackage, nil
}

func (Mpi *MembershipPackageImplementation) GetMembershipPackageByID(id uint64) (dto.MembershipPackageResponse, error) {
	var MembershipPackageModel model.MembershipPackage

	err := Mpi.db.Model(&model.MembershipPackage{}).Where("id = ? ", id).Find(&MembershipPackageModel).Error

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

func (Mpi *MembershipPackageImplementation) CreateMembershipPackage(input dto.MembershipPackageCreate) error {
	var MembershipPackageModel model.MembershipPackage

	err := copier.Copy(&MembershipPackageModel, &input)

	if err != nil {
		return err
	}

	err = Mpi.db.Model(&model.MembershipPackage{}).Create(&MembershipPackageModel).Error

	if err != nil {
		return err
	}

	return nil
}

func (Mpi *MembershipPackageImplementation) UpdateMembershipPackage(id uint64, input dto.MembershipPackageResponse) error {
	// Update menu with new data
	result := Mpi.db.Model(&model.MembershipPackage{}).Where("id = ?", id).Updates(&model.MembershipPackage{
		AdminID:  input.AdminID,
		Name:     input.Name,
		Duration: input.Duration,
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

func (Mpi *MembershipPackageImplementation) DeleteMembershipPackage(id uint64) error {
	err := Mpi.db.Where("id = ?", id).Delete(&model.MembershipPackage{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (Mpi *MembershipPackageImplementation) GetMembershipPackageByAdminID(adminID uint64) ([]dto.MembershipPackageResponse, error) {
	var MembershipPackageModel []model.MembershipPackage

	err := Mpi.db.Model(&model.MembershipPackage{}).Where("admin_id = ? ", adminID).Find(&MembershipPackageModel).Error

	// Periksa apakah FoodModel kosong (tidak ada data yang ditemukan)
	if len(MembershipPackageModel) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	var membershippackage []dto.MembershipPackageResponse

	err = copier.Copy(&membershippackage, &MembershipPackageModel)

	if err != nil {
		return nil, err
	}

	return membershippackage, nil
}

func NewMenuRepository(db *gorm.DB) MembershipPackageRepository {
	return &MembershipPackageImplementation{
		db: db,
	}
}
