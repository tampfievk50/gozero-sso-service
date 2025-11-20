package user

import (
	iport "gozero-sso-service/domain/domain-core/port/input/service"
	oport "gozero-sso-service/domain/domain-core/port/output/repository"
)

type service struct {
	userRepo oport.UserRepository
}

func NewUserService(userRepo oport.UserRepository) iport.UserService {
	return &service{
		userRepo: userRepo,
	}
}
