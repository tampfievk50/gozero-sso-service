package role

import (
	"context"
	"github.com/tampfievk50/gozero-core-api/app"
	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"github.com/zeromicro/go-zero/core/logx"
	"gozero-sso-service/internal/dataaccess/dto"
	oport "gozero-sso-service/internal/logic/port/output/repository"
)

type repository struct {
	logx.Logger
	ctx    context.Context
	svcCtx servicecontext.ServiceContextInterface
}

func (r *repository) SetState(ctx context.Context, svcCtx servicecontext.ServiceContextInterface) oport.RoleRepository {
	//TODO implement me
	panic("implement me")
}

func (r *repository) GetAllRoles() ([]dto.RoleDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) GetRole(id *uint) (*dto.RoleDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) CreateRole(RoleDto *dto.RoleDTO) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) UpdateRole(RoleDto *dto.RoleDTO) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) DeleteRole(id *uint) error {
	//TODO implement me
	panic("implement me")
}

func NewRoleRepository(ctx context.Context, svcCtx servicecontext.ServiceContextInterface) oport.RoleRepository {
	return &repository{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func init() {
	app.Bind("roleRepository", func(p *app.MakeParams) oport.RoleRepository {
		var (
			ctx    context.Context
			svcCtx servicecontext.ServiceContextInterface
		)
		for _, v := range *p {
			switch val := v.(type) {
			case context.Context:
				ctx = val
			case servicecontext.ServiceContextInterface:
				svcCtx = val
			}
		}
		return NewRoleRepository(ctx, svcCtx)
	})
}
