package repository

import (
	"context"
	"gozero-sso-service/internal/dataaccess/dto"
)

type RoleRepository interface {
	GetAllRoles(ctx context.Context) ([]dto.RoleDTO, error)
	GetRole(ctx context.Context, id *uint) (*dto.RoleDTO, error)
	CreateRole(ctx context.Context, RoleDto *dto.RoleDTO) error
	UpdateRole(ctx context.Context, RoleDto *dto.RoleDTO) error
	DeleteRole(ctx context.Context, id *uint) error
}
