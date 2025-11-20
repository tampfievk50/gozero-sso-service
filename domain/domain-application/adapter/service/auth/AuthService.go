package auth

import (
	"context"
	"gozero-sso-service/domain/domain-core/dto"
)

func (l *service) RefreshToken(ctx context.Context) (*dto.UserToken, error) {
	//TODO implement me
	panic("implement me")
}

func (l *service) Login(ctx context.Context, req *dto.UserLoginDto) (*dto.UserToken, error) {
	//TODO implement me
	panic("implement me")
}
