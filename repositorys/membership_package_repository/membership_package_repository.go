package membershippackagerepository

import "catering-api/models/dto"

type MembershipPackageRepository interface {
	GetAllMembershipPackage() ([]dto.MembershipPackageResponse, error)
	GetMembershipPackageByID(id uint64) (dto.MembershipPackageResponse, error)
	CreateMembershipPackage(input dto.MembershipPackageCreate) error
	UpdateMembershipPackage(id uint64, input dto.MembershipPackageResponse) error
	DeleteMembershipPackage(id uint64) error
	GetMembershipPackageByAdminID(adminID uint64) ([]dto.MembershipPackageResponse, error)
}
