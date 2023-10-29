package response

import (
	"catering-api/models/dto"
	"catering-api/models/model"
)

func UserDomainToUserLoginResponse(user *model.User) dto.UserResponseLogin {
	return dto.UserResponseLogin{
		Name:  user.Name,
		Email: user.Email,
		ID:    user.ID,
	}
}
