package role

import (
	"context"
	"github.com/tampfievk50/gozero-core-api/app"
	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"gozero-sso-service/internal/dataaccess/dto"
	oport "gozero-sso-service/internal/logic/port/output/repository"
)

type repository struct {
	svcCtx servicecontext.ServiceContextInterface
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

func NewRoleRepository(svcCtx servicecontext.ServiceContextInterface) oport.RoleRepository {
	return &repository{
		svcCtx: svcCtx,
	}
}

func init() {
	app.Bind("roleRepository", func(p *app.MakeParams) oport.RoleRepository {
		var (
			svcCtx servicecontext.ServiceContextInterface
		)
		for _, v := range *p {
			switch val := v.(type) {
			case servicecontext.ServiceContextInterface:
				svcCtx = val
			}
		}
		return NewRoleRepository(svcCtx)
	})
}
