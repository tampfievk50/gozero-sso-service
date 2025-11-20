package role

import (
	iport "gozero-sso-service/domain/domain-core/port/input/service"
	oport "gozero-sso-service/domain/domain-core/port/output/repository"
)

type service struct {
	roleRepo oport.RoleRepository
}

func NewRoleService(roleRepo oport.RoleRepository) iport.RoleService {
	return &service{
		roleRepo: roleRepo,
	}
}
