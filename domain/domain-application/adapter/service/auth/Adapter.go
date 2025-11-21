package auth

import (
	iport "gozero-sso-service/domain/domain-core/port/input/service"
	oport "gozero-sso-service/domain/domain-core/port/output/repository"
)

type service struct {
	userRepo oport.UserRepository
}

func NewAuthService(userRepo oport.UserRepository) iport.AuthService {
	return &service{
		userRepo: userRepo,
	}
}
