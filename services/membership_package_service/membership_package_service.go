package membershippackageservice

import (
	"catering-api/models/dto"
	membershippackagerepository "catering-api/repositorys/membership_package_repository"

	"github.com/jinzhu/copier"
)

type MembershipPackageService interface {
	GetAllMembershipPackage() ([]dto.MembershipPackageResponse, error)
	GetMembershipPackageByID(uint64) (dto.MembershipPackageResponse, error)
	CreateMembershipPackage(input dto.MembershipPackageCreate) error
	UpdateMembershipPackage(id uint64, input dto.MembershipPackageResponse) error
	DeleteMembershipPackage(id uint64) error
	GetMembershipPackageByAdminID(adminID uint64) ([]dto.MembershipPackageResponse, error)
}

type MembershipPackageImplementation struct {
	repository membershippackagerepository.MembershipPackageRepository
}

func (Mpi *MembershipPackageImplementation) GetAllMembershipPackage() ([]dto.MembershipPackageResponse, error) {
	data, err := Mpi.repository.GetAllMembershipPackage()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (Mpi *MembershipPackageImplementation) GetMembershipPackageByID(id uint64) (dto.MembershipPackageResponse, error) {
	data, err := Mpi.repository.GetMembershipPackageByID(id)
	if err != nil {
		return dto.MembershipPackageResponse{}, err
	}

	var MembershipPackage dto.MembershipPackageResponse
	err = copier.Copy(&MembershipPackage, &data)
	if err != nil {
		return dto.MembershipPackageResponse{}, err
	}

	return MembershipPackage, nil
}

func (Mpi *MembershipPackageImplementation) CreateMembershipPackage(input dto.MembershipPackageCreate) error {
	err := Mpi.repository.CreateMembershipPackage(input)
	if err != nil {
		return err
	}
	return nil
}

func (Mpi *MembershipPackageImplementation) UpdateMembershipPackage(id uint64, input dto.MembershipPackageResponse) error {
	// call repository to update course
	err := Mpi.repository.UpdateMembershipPackage(id, input)
	if err != nil {
		return err
	}
	return nil
}

func (Mpi *MembershipPackageImplementation) DeleteMembershipPackage(id uint64) error {
	err := Mpi.repository.DeleteMembershipPackage(id)
	if err != nil {
		return err
	}
	return nil
}

func (Mpi *MembershipPackageImplementation) GetMembershipPackageByAdminID(adminID uint64) ([]dto.MembershipPackageResponse, error) {
	data, err := Mpi.repository.GetMembershipPackageByAdminID(adminID)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewMenuService(MembershipPackageRepo membershippackagerepository.MembershipPackageRepository) MembershipPackageService {
	return &MembershipPackageImplementation{
		repository: MembershipPackageRepo,
	}
}
