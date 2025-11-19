package job

import (
	"context"
	"github.com/tampfievk50/gozero-core-api/app"
	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"github.com/zeromicro/go-zero/core/logx"
	oport "gozero-sso-service/internal/logic/port/output/repository"
	"gozero-sso-service/internal/svc"
)

type repository struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJobRepository(ctx context.Context, svcCtx servicecontext.ServiceContextInterface) oport.UserRepository {
	return &repository{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx.(*svc.ServiceContext),
	}
}

func init() {
	app.Bind("userRepository", func(p *app.MakeParams) oport.UserRepository {
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
		return NewJobRepository(ctx, svcCtx)
	})
}
