package permission

import (
	"context"
	"gozero-sso-service/dataaccess/model"
)

func (r *repository) DeletePermission(ctx context.Context, id string) error {
	return r.db.Connection.Model(&model.Permission{}).
		Where("id = ?", id).
		Update("is_deleted", true).Error
}
