package repository

import (
	"gozero-sso-service/internal/dataaccess/dto"
)

type RoleRepository interface {
	GetAllRoles() ([]dto.RoleDTO, error)
	GetRole(id *uint) (*dto.RoleDTO, error)
	CreateRole(RoleDto *dto.RoleDTO) error
	UpdateRole(RoleDto *dto.RoleDTO) error
	DeleteRole(id *uint) error
}
