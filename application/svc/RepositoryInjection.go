package svc

import (
	"gozero-sso-service/dataaccess/adapter/repository/role"
	"gozero-sso-service/dataaccess/adapter/repository/user"
	oport "gozero-sso-service/domain/domain-core/port/output/repository"

	"github.com/tampfievk50/gozero-core-api/ormx"
)

type Repository struct {
	UserRepository oport.UserRepository
	RoleRepository oport.RoleRepository
}

func InitRepository(database *ormx.Database) *Repository {
	return &Repository{
		UserRepository: user.NewUserRepository(database),
		RoleRepository: role.NewRoleRepository(database),
	}
}
