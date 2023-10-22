package userrepository

import (
	"catering-api/helpers"
	"catering-api/models/dto"
	"catering-api/models/model"
	"errors"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UserImplementation struct {
	db *gorm.DB
}

func (Ui *UserImplementation) GetAllUser() ([]dto.UserResponse,error)  {
	var UserModel []model.User

	err := Ui.db.Find(&UserModel).Error

	if err != nil {
		return nil, err
	}

	var User []dto.UserResponse

	if err := copier.Copy(&User, &UserModel); err != nil {
		return nil, err
	}

	return User , nil
}

func(Ui *UserImplementation) GetUserByID(id uint64) (dto.UserResponse , error)  {
	var UserModel model.User

	err := Ui.db.Model(&model.User{}).Where("id = ? ", id).Find(&UserModel).Error

	if err != nil {
		return dto.UserResponse{}, err
	}

	// Periksa apakah data ditemukan
	if UserModel.ID == 0 {
		return dto.UserResponse{}, gorm.ErrRecordNotFound
	}
	
	var User dto.UserResponse

	err = copier.Copy(&User,&UserModel)

	if err != nil {
		return dto.UserResponse{}, err
	}

	return User, nil 
}

func (Ui *UserImplementation) CreateUser(input dto.UserCreate) error {
    var UserModel model.User

    err := copier.Copy(&UserModel, &input)
    if err != nil {
        return err
    }

    // Cek apakah email sudah ada di database
    var emailCount int64
    Ui.db.Model(&model.User{}).Where("email = ?", input.Email).Count(&emailCount)
    if emailCount > 0 {
        return errors.New("email already use")
    }

    // Tambahkan pengguna ke database jika email dan username belum ada
    err = Ui.db.Model(&model.User{}).Create(&UserModel).Error
    if err != nil {
        return err
    }

    return nil
}


func (Ui *UserImplementation) UpdateUser(id uint64, input dto.UserCreate) error {
	// Update menu with new data
	result := Ui.db.Model(&model.User{}).Where("id = ?", id).Updates(&model.User{
		Name:     input.Name,
		Email: input.Email,
		Password: input.Password,
		Address: input.Address,
		Phone: input.Phone,
		Gender: input.Gender,
		MembershipType: input.MembershipType,
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


func (Ui *UserImplementation) DeleteUser(id uint64) error {
	err := Ui.db.Where("id = ?", id).Delete(&model.User{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (Ui *UserImplementation) LoginUser(input dto.UserLogin) (dto.UserResponse, error) {
	var UserLogin dto.UserResponse

	err := Ui.db.Model(&model.User{}).First(&UserLogin, "email = ?", input.Email).Error

	if err != nil {
		return dto.UserResponse{}, errors.New("email not registered")
	}

	match := helpers.CheckPasswordHash(input.Password, UserLogin.Password)

	if !match {
		return dto.UserResponse{}, errors.New("wrong password")
	}
	var UserLoginResponse = dto.UserResponse{
		ID:           input.ID,
		Email:        input.Email,
		Password:     input.Password,
	}
	
	return UserLoginResponse, nil
}

func NewUserRepository(db *gorm.DB) UserRepository  {
	return &UserImplementation{
		db: db,
	}
}