package service

import (
	"context"
	"gozero-sso-service/domain/domain-core/dto"
)

type RoleService interface {
	GetAllRoles(ctx context.Context) ([]dto.RoleDTO, error)
	GetRole(ctx context.Context, id uint) (*dto.RoleDTO, error)
	CreateRole(ctx context.Context, createRoleRequest *dto.RoleDTO) error
	UpdateRole(ctx context.Context, roleDto *dto.RoleDTO) error
	DeleteRole(ctx context.Context, id uint) error
	AddPolicy(ctx context.Context, rolePermissionDto *dto.RolePermissionDto) error
}
