package role

import (
	"context"

	"gozero-sso-service/dataaccess/model"
)

func (r *repository) DeleteRole(ctx context.Context, id string) error {
	return r.db.Connection.Model(&model.Role{}).
		Where("id = ?", id).
		Update("is_deleted", true).Error
}
