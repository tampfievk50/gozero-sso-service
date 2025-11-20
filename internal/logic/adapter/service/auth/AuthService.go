package user

import (
	"context"
	"gozero-sso-service/internal/dataaccess/dto"
	"gozero-sso-service/internal/types"
)

func (l *service) RefreshToken(ctx context.Context) (*dto.UserToken, error) {
	//TODO implement me
	panic("implement me")
}

func (l *service) Login(ctx context.Context, req *types.UserLoginRequest) (*dto.UserToken, error) {
	//TODO implement me
	panic("implement me")
}
