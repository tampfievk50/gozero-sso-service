package svc

import (
	"github.com/casbin/casbin/v2"
	"gozero-sso-service/domain/domain-application/adapter/service"
	"gozero-sso-service/domain/domain-application/adapter/service/auth"
	"gozero-sso-service/domain/domain-application/adapter/service/role"
	"gozero-sso-service/domain/domain-application/adapter/service/user"
	"gozero-sso-service/domain/domain-core/port/output/repository"
)

func InitService(rp *repository.Repository, enforcer *casbin.SyncedEnforcer) *service.Service {
	return &service.Service{
		AuthService: auth.NewAuthService(rp, enforcer),
		UserService: user.NewUserService(rp, enforcer),
		RoleService: role.NewRoleService(rp, enforcer),
	}
}
