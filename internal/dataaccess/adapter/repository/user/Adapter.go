package job

import (
	"github.com/tampfievk50/gozero-core-api/app"
	"github.com/tampfievk50/gozero-core-api/servicecontext"
	oport "gozero-sso-service/internal/logic/port/output/repository"
	"gozero-sso-service/internal/svc"
)

type repository struct {
	svcCtx *svc.ServiceContext
}

func NewJobRepository(svcCtx servicecontext.ServiceContextInterface) oport.UserRepository {
	return &repository{
		svcCtx: svcCtx.(*svc.ServiceContext),
	}
}

func init() {
	app.Bind("userRepository", func(p *app.MakeParams) oport.UserRepository {
		var (
			svcCtx servicecontext.ServiceContextInterface
		)
		for _, v := range *p {
			switch val := v.(type) {
			case servicecontext.ServiceContextInterface:
				svcCtx = val
			}
		}
		return NewJobRepository(svcCtx)
	})
}
