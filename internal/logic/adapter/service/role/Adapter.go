package Role

import (
	"context"
	"github.com/tampfievk50/gozero-core-api/app"
	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"github.com/zeromicro/go-zero/core/logx"
	iport "gozero-sso-service/internal/logic/port/input/service"
	"gozero-sso-service/internal/svc"
)

type service struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleService(ctx context.Context, svcCtx servicecontext.ServiceContextInterface) iport.RoleService {
	return &service{
		ctx:    ctx,
		svcCtx: svcCtx.(*svc.ServiceContext),
		Logger: logx.WithContext(ctx),
	}
}

func init() {
	app.Bind("roleService", func(p *app.MakeParams) iport.RoleService {
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
		return NewRoleService(ctx, svcCtx)
	})
}
