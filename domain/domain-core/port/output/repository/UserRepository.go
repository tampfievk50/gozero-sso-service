package repository

import (
	"context"
	"gozero-sso-service/domain/domain-core/dto"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context, pager *dto.PagerDto) ([]dto.UserDTO, error)
	GetUser(ctx context.Context, id string) (*dto.UserDTO, error)
	CreateUser(ctx context.Context, userDto *dto.UserDTO) error
	UpdateUser(ctx context.Context, userDto *dto.UserDTO) error
	DeleteUser(ctx context.Context, id string) error
	GetUserByMail(ctx context.Context, email *string) (*dto.UserDTO, error)
	GetUserByUsername(ctx context.Context, username string) (*dto.UserDTO, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	ExistsByUsername(ctx context.Context, username string) (bool, error)
	GetUserRoles(ctx context.Context, userId string) ([]dto.RoleDTO, error)
	AssignRoles(ctx context.Context, userId string, roleIDs []string, domainID string) error
	RemoveRole(ctx context.Context, userId string, roleID string, domainID string) error
}
