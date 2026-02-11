package service

import (
	"context"
	"gozero-sso-service/domain/domain-core/dto"
)

type RoleService interface {
	GetAllRoles(ctx context.Context, pager *dto.PagerDto) ([]dto.RoleDTO, error)
	GetRole(ctx context.Context, id string) (*dto.RoleDTO, error)
	CreateRole(ctx context.Context, createRoleRequest *dto.RoleDTO) error
	UpdateRole(ctx context.Context, roleDto *dto.RoleDTO) error
	DeleteRole(ctx context.Context, id string) error
	AddPolicy(ctx context.Context, rolePermissionDto *dto.RolePermissionDto) error
	GetAllPolicies(ctx context.Context, roleID string) ([][]string, error)
	GetRolePolicies(ctx context.Context, roleID string) ([][]string, error)
	RemovePolicy(ctx context.Context, roleID string, domainID string, path string, action string) error
}
