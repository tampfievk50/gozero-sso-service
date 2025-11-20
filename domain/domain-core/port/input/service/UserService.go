package service

import (
	"context"
	"gozero-sso-service/domain/domain-core/dto"
)

type UserService interface {
	GetAllUsers(ctx context.Context) ([]dto.UserDTO, error)
	GetUser(ctx context.Context, id uint) (*dto.UserDTO, error)
	CreateUser(ctx context.Context, createUserRequest *dto.UserDTO) error
	UpdateUser(ctx context.Context, userDto *dto.UserDTO) error
	DeleteUser(ctx context.Context, id uint) error
}
