package svc

import (
	"gozero-sso-service/application/core"
	"gozero-sso-service/domain/domain-application/adapter/service/auth"
	"gozero-sso-service/domain/domain-application/adapter/service/role"
	"gozero-sso-service/domain/domain-application/adapter/service/user"
)

func InitService(rp *core.Repository) *core.Service {
	return &core.Service{
		AuthService: auth.NewAuthService(rp.UserRepository),
		UserService: user.NewUserService(rp.UserRepository),
		RoleService: role.NewRoleService(rp.RoleRepository),
	}
}
