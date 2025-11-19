package service

import (
	"gozero-sso-service/internal/dataaccess/dto"
	"gozero-sso-service/internal/types"
)

type AuthService interface {
	RefreshToken() (*dto.UserToken, error)
	Login(req *types.UserLoginRequest) (*dto.UserToken, error)
}
