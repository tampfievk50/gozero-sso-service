package role

import (
	"context"
	"gozero-sso-service/domain/domain-core/dto"
	oport "gozero-sso-service/domain/domain-core/port/output/repository"

	"github.com/tampfievk50/gozero-core-api/ormx"
)

type repository struct {
	db *ormx.Database
}

func (r *repository) GetAllRoles(ctx context.Context) ([]dto.RoleDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) GetRole(ctx context.Context, id *uint) (*dto.RoleDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) CreateRole(ctx context.Context, RoleDto *dto.RoleDTO) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) UpdateRole(ctx context.Context, RoleDto *dto.RoleDTO) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) DeleteRole(ctx context.Context, id *uint) error {
	//TODO implement me
	panic("implement me")
}

func NewRoleRepository(db *ormx.Database) oport.RoleRepository {
	return &repository{
		db: db,
	}
}
