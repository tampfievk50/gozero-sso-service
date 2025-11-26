package auth

import (
	"gozero-sso-service/dataaccess/adapter/repository"
	iport "gozero-sso-service/domain/domain-core/port/input/service"
)

type service struct {
	rp *repository.Repository
}

func NewAuthService(rp *repository.Repository) iport.AuthService {
	return &service{
		rp: rp,
	}
}
