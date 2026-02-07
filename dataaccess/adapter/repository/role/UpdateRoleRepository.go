package role

import (
	"context"

	"gozero-sso-service/dataaccess/model"
	"gozero-sso-service/domain/domain-core/dto"
)

func (r *repository) UpdateRole(ctx context.Context, roleDto *dto.RoleDTO) error {
	var role model.Role

	err := r.db.Connection.First(&role, roleDto.ID).Error
	if err != nil {
		return err
	}

	if roleDto.Name != "" {
		role.Name = roleDto.Name
	}
	if roleDto.Description != "" {
		role.Description = roleDto.Description
	}

	return r.db.Connection.Save(&role).Error
}
