package user

import (
	oport "gozero-sso-service/domain/domain-core/port/output/repository"

	"github.com/tampfievk50/gozero-core-api/ormx"
)

type repository struct {
	db *ormx.Database
}

func NewUserRepository(db *ormx.Database) oport.UserRepository {
	return &repository{
		db: db,
	}
}
