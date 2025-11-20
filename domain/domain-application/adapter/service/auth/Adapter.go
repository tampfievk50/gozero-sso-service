package auth

import (
	iport "gozero-sso-service/domain/domain-core/port/input/service"
)

type service struct {
	userService iport.UserService
}

func NewAuthService(userService iport.UserService) iport.AuthService {
	return &service{
		userService: userService,
	}
}
