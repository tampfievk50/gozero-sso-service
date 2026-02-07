package service

import (
	"context"
	"gozero-sso-service/domain/domain-core/dto"
)

type UserService interface {
	GetAllUsers(ctx context.Context, pager *dto.PagerDto) ([]dto.UserDTO, error)
	GetUser(ctx context.Context, id uint) (*dto.UserDTO, error)
	CreateUser(ctx context.Context, createUserRequest *dto.UserDTO) error
	UpdateUser(ctx context.Context, userDto *dto.UserDTO) error
	DeleteUser(ctx context.Context, id uint) error
	GetUserRoles(ctx context.Context, userId uint) ([]dto.RoleDTO, error)
	AssignRoles(ctx context.Context, userId uint, roleIDs []uint, domainID uint) error
	RemoveRole(ctx context.Context, userId uint, roleID uint, domainID uint) error
	GetUserPermissions(ctx context.Context, userId uint) ([]dto.PermissionDTO, error)
}
