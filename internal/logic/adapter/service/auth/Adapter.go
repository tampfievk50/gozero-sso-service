package user

import (
	"github.com/tampfievk50/gozero-core-api/app"
	"github.com/tampfievk50/gozero-core-api/servicecontext"
	iport "gozero-sso-service/internal/logic/port/input/service"
	"gozero-sso-service/internal/svc"
)

type service struct {
	svcCtx *svc.ServiceContext
}

func NewAuthService(svcCtx servicecontext.ServiceContextInterface) iport.AuthService {
	return &service{
		svcCtx: svcCtx.(*svc.ServiceContext),
	}
}

func init() {
	app.Bind("authService", func(p *app.MakeParams) iport.AuthService {
		var (
			svcCtx servicecontext.ServiceContextInterface
		)
		for _, v := range *p {
			switch val := v.(type) {
			case servicecontext.ServiceContextInterface:
				svcCtx = val
			}
		}
		return NewAuthService(svcCtx)
	})
}
