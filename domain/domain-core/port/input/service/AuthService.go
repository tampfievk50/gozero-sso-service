package service

import (
	"context"
	"github.com/casbin/casbin/v2"
	"gozero-sso-service/domain/domain-core/dto"
)

type AuthService interface {
	RefreshToken(ctx context.Context) (*dto.UserToken, error)
	Login(ctx context.Context, req *dto.UserLoginDto, config *dto.ConfigDto) (*dto.UserToken, error)
	HasPermission(ctx context.Context, m *casbin.SyncedEnforcer, subs []string, doms []string, path, method string) (bool, error)
}
