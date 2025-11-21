package svc

import (
	"gozero-sso-service/domain/domain-application/adapter/service/auth"
	"gozero-sso-service/domain/domain-application/adapter/service/role"
	"gozero-sso-service/domain/domain-application/adapter/service/user"
	iport "gozero-sso-service/domain/domain-core/port/input/service"
)

type Service struct {
	AuthService iport.AuthService
	UserService iport.UserService
	RoleService iport.RoleService
}

func InitService(rp *Repository) *Service {
	return &Service{
		AuthService: auth.NewAuthService(rp.UserRepository),
		UserService: user.NewUserService(rp.UserRepository),
		RoleService: role.NewRoleService(rp.RoleRepository),
	}
}
