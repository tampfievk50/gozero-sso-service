package svc

import (
	"gozero-sso-service/dataaccess/adapter/repository"
	"gozero-sso-service/dataaccess/adapter/repository/role"
	"gozero-sso-service/dataaccess/adapter/repository/user"

	"github.com/tampfievk50/gozero-core-api/ormx"
)

func InitRepository(database *ormx.Database) *repository.Repository {
	return &repository.Repository{
		UserRepository: user.NewUserRepository(database),
		RoleRepository: role.NewRoleRepository(database),
	}
}
