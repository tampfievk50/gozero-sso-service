package repository

import (
	"context"
	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"gozero-sso-service/internal/dataaccess/dto"
)

type RoleRepository interface {
	SetState(ctx context.Context, svcCtx servicecontext.ServiceContextInterface) RoleRepository
	GetAllRoles() ([]dto.RoleDTO, error)
	GetRole(id *uint) (*dto.RoleDTO, error)
	CreateRole(RoleDto *dto.RoleDTO) error
	UpdateRole(RoleDto *dto.RoleDTO) error
	DeleteRole(id *uint) error
}
