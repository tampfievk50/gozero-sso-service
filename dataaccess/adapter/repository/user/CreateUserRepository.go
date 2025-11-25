package user

import (
	"context"
	"gozero-sso-service/dataaccess/model"
	"gozero-sso-service/domain/domain-core/dto"

	"github.com/jinzhu/copier"
)

func (r *repository) CreateUser(ctx context.Context, userDto *dto.UserDTO) error {
	var (
		user model.User
	)

	err := copier.Copy(&user, userDto)
	if err != nil {
		return err
	}

	for _, id := range userDto.RoleIDs {
		user.Roles = append(user.Roles, model.Role{VModel: model.VModel{
			ID: id,
		}})
	}

	for _, id := range userDto.ResourceIDs {
		user.Resources = append(user.Resources, model.Resource{VModel: model.VModel{
			ID: id,
		}})
	}

	if err = r.db.Connection.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
