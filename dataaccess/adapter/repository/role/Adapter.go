package role

import (
	oport "gozero-sso-service/domain/domain-core/port/output/repository"

	"github.com/tampfievk50/gozero-core-api/ormx"
)

type repository struct {
	db *ormx.Database
}

func NewRoleRepository(db *ormx.Database) oport.RoleRepository {
	return &repository{
		db: db,
	}
}
