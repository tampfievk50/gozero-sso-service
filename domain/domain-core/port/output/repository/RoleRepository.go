package repository

import (
	"context"
	"gozero-sso-service/domain/domain-core/dto"
)

type RoleRepository interface {
	GetAllRoles(ctx context.Context, pager *dto.PagerDto) ([]dto.RoleDTO, error)
	GetRole(ctx context.Context, id uint) (*dto.RoleDTO, error)
	GetRoleByName(ctx context.Context, name string) (*dto.RoleDTO, error)
	ByIDs(ctx context.Context, ids []uint) ([]dto.RoleDTO, error)
	CreateRole(ctx context.Context, RoleDto *dto.RoleDTO) error
	UpdateRole(ctx context.Context, RoleDto *dto.RoleDTO) error
	DeleteRole(ctx context.Context, id uint) error
}
