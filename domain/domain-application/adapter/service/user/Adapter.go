package user

import (
	"gozero-sso-service/dataaccess/adapter/repository"
	iport "gozero-sso-service/domain/domain-core/port/input/service"
)

type service struct {
	rp *repository.Repository
}

func NewUserService(rp *repository.Repository) iport.UserService {
	return &service{
		rp: rp,
	}
}
