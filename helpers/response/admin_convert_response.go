package response

import (
	"catering-api/models/dto"
	"catering-api/models/model"
)

func AdminModelToAdminLoginResponse(admin *model.Admin) dto.AdminResponseLogin {
	return dto.AdminResponseLogin{
		Name:  admin.Name,
		Email: admin.Email,
		ID:    admin.ID,
	}
}
