package Role

import (
	"github.com/tampfievk50/gozero-core-api/app"
	"github.com/tampfievk50/gozero-core-api/servicecontext"
	iport "gozero-sso-service/internal/logic/port/input/service"
	"gozero-sso-service/internal/svc"
)

type service struct {
	svcCtx *svc.ServiceContext
}

func NewRoleService(svcCtx servicecontext.ServiceContextInterface) iport.RoleService {
	return &service{
		svcCtx: svcCtx.(*svc.ServiceContext),
	}
}

func init() {
	app.Bind("roleService", func(p *app.MakeParams) iport.RoleService {
		var (
			svcCtx servicecontext.ServiceContextInterface
		)
		for _, v := range *p {
			switch val := v.(type) {
			case servicecontext.ServiceContextInterface:
				svcCtx = val
			}
		}
		return NewRoleService(svcCtx)
	})
}
