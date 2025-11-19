package service

import (
	"gozero-sso-service/internal/dataaccess/dto"
	"gozero-sso-service/internal/types"
)

type RoleService interface {
	GetAllRoles() ([]dto.RoleDTO, error)
	GetRole(req *types.IdPathRequest) (*dto.RoleDTO, error)
	CreateRole(createRoleRequest *types.CreateRoleRequest) error
	UpdateRole(RoleDto *types.UpdateRoleRequest) error
	DeleteRole(req *types.IdPathRequest) error
}
