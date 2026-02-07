package permission

import (
	oport "gozero-sso-service/domain/domain-core/port/output/repository"
	"github.com/tampfievk50/gozero-core-api/ormx"
)

type repository struct {
	db *ormx.Database
}

func NewPermissionRepository(db *ormx.Database) oport.PermissionRepository {
	return &repository{
		db: db,
	}
}
