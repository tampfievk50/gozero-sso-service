package user

import (
	"context"
	"gozero-sso-service/dataaccess/model"
	"gozero-sso-service/domain/domain-core/dto"
)

func (r *repository) UpdateUser(ctx context.Context, userDto *dto.UserDTO) error {
	var user model.User

	err := r.db.Connection.First(&user, userDto.ID).Error
	if err != nil {
		return err
	}
	if userDto.FirstName != nil {
		user.FirstName = *userDto.FirstName
	}
	if userDto.LastName != nil {
		user.LastName = *userDto.LastName
	}
	if userDto.DateOfBirth != nil {
		user.DateOfBirth = userDto.DateOfBirth
	}
	if userDto.Gender != nil {
		user.Gender = *userDto.Gender
	}
	if userDto.IsActive != nil {
		user.IsActive = *userDto.IsActive
	}
	err = r.db.Connection.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}
