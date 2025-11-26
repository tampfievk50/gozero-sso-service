package svc

import (
	"gozero-sso-service/dataaccess/adapter/repository"
	"gozero-sso-service/domain/domain-application/adapter/service"
	"gozero-sso-service/domain/domain-application/adapter/service/auth"
	"gozero-sso-service/domain/domain-application/adapter/service/role"
	"gozero-sso-service/domain/domain-application/adapter/service/user"
)

func InitService(rp *repository.Repository) *service.Service {
	return &service.Service{
		AuthService: auth.NewAuthService(rp),
		UserService: user.NewUserService(rp),
		RoleService: role.NewRoleService(rp),
	}
}
