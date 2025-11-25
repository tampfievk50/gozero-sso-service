package service

import (
	"context"
	"gozero-sso-service/domain/domain-core/dto"

	"github.com/casbin/casbin/v2"
)

type AuthService interface {
	RefreshToken(ctx context.Context) (*dto.UserToken, error)
	Login(ctx context.Context, req *dto.UserLoginDto, config *dto.ConfigDto) (*dto.UserToken, error)
	HasPermission(ctx context.Context, m *casbin.SyncedEnforcer, sub string, doms []string, path, method string) (bool, error)
}
