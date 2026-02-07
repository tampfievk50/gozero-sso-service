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

	if err = r.db.Connection.Create(&user).Error; err != nil {
		return err
	}
	userDto.ID = user.ID
	return nil
}
