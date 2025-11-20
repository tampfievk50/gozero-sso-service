package service

import (
	"context"
	"gozero-sso-service/domain/domain-core/dto"
)

type AuthService interface {
	RefreshToken(ctx context.Context) (*dto.UserToken, error)
	Login(ctx context.Context, req *dto.UserLoginDto) (*dto.UserToken, error)
}
