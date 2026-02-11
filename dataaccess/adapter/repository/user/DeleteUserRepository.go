package user

import (
	"context"
	"errors"
	"gozero-sso-service/dataaccess/model"
)

func (r *repository) DeleteUser(ctx context.Context, id string) error {
	var user model.User
	err := r.db.Connection.First(&user, "id = ?", id).Error
	if err != nil {
		return err
	}
	if user.IsDeleted || user.IsSupper {
		return errors.New("user cannot be deleted")
	}
	return r.db.Connection.Model(&model.User{}).
		Where("id = ?", id).
		Update("is_deleted", true).
		Update("is_active", false).Error
}
