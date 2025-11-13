package repository

import (
	"context"
	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"gozero-sso-service/internal/dataaccess/dto"
)

type UserRepository interface {
	SetState(ctx context.Context, svcCtx servicecontext.ServiceContextInterface) UserRepository
	GetAllUsers() ([]dto.UserDTO, error)
	GetUser(id *uint) (*dto.UserDTO, error)
	CreateUser(userDto *dto.UserDTO) error
	UpdateUser(userDto *dto.UserDTO) error
	DeleteUser(id *uint) error
}
