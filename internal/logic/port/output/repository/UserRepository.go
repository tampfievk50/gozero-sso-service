package repository

import (
	"context"
	"gozero-sso-service/internal/dataaccess/dto"
	"gozero-sso-service/internal/types"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context, pager *types.Pager) ([]dto.UserDTO, error)
	GetUser(ctx context.Context, id *uint) (*dto.UserDTO, error)
	CreateUser(ctx context.Context, userDto *dto.UserDTO) error
	UpdateUser(ctx context.Context, userDto *dto.UserDTO) error
	DeleteUser(ctx context.Context, id *uint) error
}
