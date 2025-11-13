package service

import (
	"context"
	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"gozero-sso-service/internal/dataaccess/dto"
	"gozero-sso-service/internal/types"
)

type UserService interface {
	SetState(ctx context.Context, svcCtx servicecontext.ServiceContextInterface) UserService
	GetAllUsers() ([]dto.UserDTO, error)
	GetUser(req *types.IdPathRequest) (*dto.UserDTO, error)
	CreateUser(createUserRequest *types.CreateUserRequest) error
	UpdateUser(userDto *types.UpdateUserRequest) error
	DeleteUser(req *types.IdPathRequest) error
}
