package svc

import (
	"gozero-sso-service/dataaccess/adapter/repository"
	"gozero-sso-service/domain/domain-application/adapter/service"
	"gozero-sso-service/domain/domain-application/adapter/service/auth"
	"gozero-sso-service/domain/domain-application/adapter/service/role"
	"gozero-sso-service/domain/domain-application/adapter/service/user"

	"github.com/casbin/casbin/v2"
)

func InitService(rp *repository.Repository, m *casbin.SyncedEnforcer) *service.Service {
	return &service.Service{
		AuthService: auth.NewAuthService(rp.UserRepository, m),
		UserService: user.NewUserService(rp.UserRepository),
		RoleService: role.NewRoleService(rp.RoleRepository),
	}
}
