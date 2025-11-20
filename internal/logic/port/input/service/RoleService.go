package service

import (
	"context"
	"gozero-sso-service/internal/dataaccess/dto"
	"gozero-sso-service/internal/types"
)

type RoleService interface {
	GetAllRoles(ctx context.Context) ([]dto.RoleDTO, error)
	GetRole(ctx context.Context, req *types.IdPathRequest) (*dto.RoleDTO, error)
	CreateRole(ctx context.Context, createRoleRequest *types.CreateRoleRequest) error
	UpdateRole(ctx context.Context, RoleDto *types.UpdateRoleRequest) error
	DeleteRole(ctx context.Context, req *types.IdPathRequest) error
}
