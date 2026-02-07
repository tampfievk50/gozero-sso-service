package permission

import (
	"context"
	"gozero-sso-service/domain/domain-core/dto"
)

func (l *service) GetAllPermissions(ctx context.Context, pager *dto.PagerDto) ([]dto.PermissionDTO, error) {
	return l.rp.PermissionRepository.GetAllPermissions(ctx, pager)
}

func (l *service) GetPermission(ctx context.Context, id uint) (*dto.PermissionDTO, error) {
	return l.rp.PermissionRepository.GetPermission(ctx, id)
}
