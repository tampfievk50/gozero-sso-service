package service

import (
	"context"
	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"gozero-sso-service/internal/dataaccess/dto"
	"gozero-sso-service/internal/types"
)

type AuthService interface {
	SetState(ctx context.Context, svcCtx servicecontext.ServiceContextInterface) AuthService
	RefreshToken() (*dto.UserToken, error)
	Login(req *types.UserLoginRequest) (*dto.UserToken, error)
}
