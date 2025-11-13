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

func NewJobRepository() oport.UserRepository {
	return &repository{}
}

func (r *repository) SetState(ctx context.Context, svcCtx servicecontext.ServiceContextInterface) oport.UserRepository {
	r.ctx = ctx
	r.svcCtx = svcCtx.(*svc.ServiceContext)
	return r
}

func init() {
	app.Register("userRepository", func() oport.UserRepository {
		return NewJobRepository()
	})
}
