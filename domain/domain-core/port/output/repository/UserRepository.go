package repository

import (
	"context"
	"gozero-sso-service/domain/domain-core/dto"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context, pager *dto.PagerDto) ([]dto.UserDTO, error)
	GetUser(ctx context.Context, id *uint) (*dto.UserDTO, error)
	CreateUser(ctx context.Context, userDto *dto.UserDTO) error
	UpdateUser(ctx context.Context, userDto *dto.UserDTO) error
	DeleteUser(ctx context.Context, id *uint) error
	GetUserByMail(ctx context.Context, email *string) (*dto.UserDTO, error)
}
