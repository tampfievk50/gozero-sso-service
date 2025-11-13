package service

import (
	"context"
	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"gozero-sso-service/internal/dataaccess/dto"
	"gozero-sso-service/internal/types"
)

type RoleService interface {
	SetState(ctx context.Context, svcCtx servicecontext.ServiceContextInterface) RoleService
	GetAllRoles() ([]dto.RoleDTO, error)
	GetRole(req *types.IdPathRequest) (*dto.RoleDTO, error)
	CreateRole(createRoleRequest *types.CreateRoleRequest) error
	UpdateRole(RoleDto *types.UpdateRoleRequest) error
	DeleteRole(req *types.IdPathRequest) error
}
