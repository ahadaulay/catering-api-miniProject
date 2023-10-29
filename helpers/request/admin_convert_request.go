package request

import (
	"catering-api/models/dto"
	"catering-api/models/model"
)

func AdminLoginRequestToAdminDomain(request dto.AdminLogin) *model.Admin {
	return &model.Admin{
		Email:    request.Email,
		ID:       request.ID,
		Password: request.Password,
	}
}
