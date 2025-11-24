package svc

import (
	"gozero-sso-service/application/core"
	"gozero-sso-service/dataaccess/adapter/repository/role"
	"gozero-sso-service/dataaccess/adapter/repository/user"

	"github.com/tampfievk50/gozero-core-api/ormx"
)

func InitRepository(database *ormx.Database) *core.Repository {
	return &core.Repository{
		UserRepository: user.NewUserRepository(database),
		RoleRepository: role.NewRoleRepository(database),
	}
}
