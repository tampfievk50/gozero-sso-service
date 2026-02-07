package svc

import (
	"gozero-sso-service/dataaccess/adapter/repository/permission"
	"gozero-sso-service/dataaccess/adapter/repository/role"
	"gozero-sso-service/dataaccess/adapter/repository/user"
	"gozero-sso-service/domain/domain-core/port/output/repository"

	"github.com/tampfievk50/gozero-core-api/ormx"
)

func InitRepository(database *ormx.Database) *repository.Repository {
	return &repository.Repository{
		UserRepository:       user.NewUserRepository(database),
		RoleRepository:       role.NewRoleRepository(database),
		PermissionRepository: permission.NewPermissionRepository(database),
	}
}
