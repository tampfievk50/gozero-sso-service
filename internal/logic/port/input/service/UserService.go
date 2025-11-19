package service

import (
	"gozero-sso-service/internal/dataaccess/dto"
	"gozero-sso-service/internal/types"
)

type UserService interface {
	GetAllUsers() ([]dto.UserDTO, error)
	GetUser(req *types.IdPathRequest) (*dto.UserDTO, error)
	CreateUser(createUserRequest *types.CreateUserRequest) error
	UpdateUser(userDto *types.UpdateUserRequest) error
	DeleteUser(req *types.IdPathRequest) error
}
