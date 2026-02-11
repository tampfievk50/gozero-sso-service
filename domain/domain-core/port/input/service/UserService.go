package service

import (
	"context"
	"gozero-sso-service/domain/domain-core/dto"
)

type UserService interface {
	GetAllUsers(ctx context.Context, pager *dto.PagerDto) ([]dto.UserDTO, error)
	GetUser(ctx context.Context, id string) (*dto.UserDTO, error)
	CreateUser(ctx context.Context, createUserRequest *dto.UserDTO) error
	UpdateUser(ctx context.Context, userDto *dto.UserDTO) error
	DeleteUser(ctx context.Context, id string) error
	GetUserRoles(ctx context.Context, userId string) ([]dto.RoleDTO, error)
	AssignRoles(ctx context.Context, userId string, roleIDs []string, domainID string) error
	RemoveRole(ctx context.Context, userId string, roleID string, domainID string) error
	GetUserPermissions(ctx context.Context, userId string) ([]dto.PermissionDTO, error)
}
