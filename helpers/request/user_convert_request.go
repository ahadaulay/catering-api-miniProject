package request

import (
	"catering-api/models/dto"
	"catering-api/models/model"
)

func UserLoginRequestToUserDomain(request dto.UserLogin) *model.User {
	return &model.User{
		Email:    request.Email,
		Password: request.Password,
	}
}
