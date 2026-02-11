package permission

import (
	"context"
	"gozero-sso-service/dataaccess/model"
	"gozero-sso-service/domain/domain-core/dto"
)

func (r *repository) UpdatePermission(ctx context.Context, permissionDto *dto.PermissionDTO) error {
	var permission model.Permission

	err := r.db.Connection.First(&permission, "id = ?", permissionDto.ID).Error
	if err != nil {
		return err
	}

	if permissionDto.Name != "" {
		permission.Name = permissionDto.Name
	}
	if permissionDto.Description != "" {
		permission.Description = permissionDto.Description
	}
	if permissionDto.Path != "" {
		permission.Path = permissionDto.Path
	}
	if permissionDto.Action != "" {
		permission.Action = permissionDto.Action
	}

	return r.db.Connection.Save(&permission).Error
}
