package service

import (
	"context"
	"gozero-sso-service/internal/dataaccess/dto"
	"gozero-sso-service/internal/types"
)

type UserService interface {
	GetAllUsers(ctx context.Context) ([]dto.UserDTO, error)
	GetUser(ctx context.Context, req *types.IdPathRequest) (*dto.UserDTO, error)
	CreateUser(ctx context.Context, createUserRequest *types.CreateUserRequest) error
	UpdateUser(ctx context.Context, userDto *types.UpdateUserRequest) error
	DeleteUser(ctx context.Context, req *types.IdPathRequest) error
}
