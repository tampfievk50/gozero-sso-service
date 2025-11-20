package service

import (
	"context"
	"gozero-sso-service/internal/dataaccess/dto"
	"gozero-sso-service/internal/types"
)

type AuthService interface {
	RefreshToken(ctx context.Context) (*dto.UserToken, error)
	Login(ctx context.Context, req *types.UserLoginRequest) (*dto.UserToken, error)
}
