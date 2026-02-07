package service

import (
	"context"
	"gozero-sso-service/domain/domain-core/dto"
)

type PermissionService interface {
	GetAllPermissions(ctx context.Context, pager *dto.PagerDto) ([]dto.PermissionDTO, error)
	GetPermission(ctx context.Context, id uint) (*dto.PermissionDTO, error)
	CreatePermission(ctx context.Context, permissionDto *dto.PermissionDTO) error
	UpdatePermission(ctx context.Context, permissionDto *dto.PermissionDTO) error
	DeletePermission(ctx context.Context, id uint) error
}
