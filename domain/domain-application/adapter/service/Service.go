package service

import iport "gozero-sso-service/domain/domain-core/port/input/service"

type Service struct {
	AuthService iport.AuthService
	UserService iport.UserService
	RoleService iport.RoleService
}
