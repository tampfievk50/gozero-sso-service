package permission

import (
	"context"
	"gozero-sso-service/dataaccess/model"
	"gozero-sso-service/domain/domain-core/dto"

	"github.com/jinzhu/copier"
)

func (r *repository) CreatePermission(ctx context.Context, permissionDto *dto.PermissionDTO) error {
	var permission model.Permission

	err := copier.Copy(&permission, permissionDto)
	if err != nil {
		return err
	}

	return r.db.Connection.Create(&permission).Error
}
