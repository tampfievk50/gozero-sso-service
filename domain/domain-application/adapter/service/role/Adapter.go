package role

import (
	"gozero-sso-service/dataaccess/adapter/repository"
	iport "gozero-sso-service/domain/domain-core/port/input/service"
)

type service struct {
	rp *repository.Repository
}

func NewRoleService(rp *repository.Repository) iport.RoleService {
	return &service{
		rp: rp,
	}
}
