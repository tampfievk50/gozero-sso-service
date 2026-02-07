package permission

import (
	"context"
	"gozero-sso-service/domain/domain-core/dto"
)

func (l *service) UpdatePermission(ctx context.Context, permissionDto *dto.PermissionDTO) error {
	return l.rp.PermissionRepository.UpdatePermission(ctx, permissionDto)
}
