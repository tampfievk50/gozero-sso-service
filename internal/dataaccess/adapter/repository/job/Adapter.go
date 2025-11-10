package job

import (
	"context"
	"github.com/tampfievk50/gozero-core-api/app"
	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"github.com/zeromicro/go-zero/core/logx"
	op "gozero-sso-service/internal/logic/ports/output/repository"
	"gozero-sso-service/internal/svc"
)

type repository struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJobRepository() op.JobRepository {
	return &repository{}
}

func (r *repository) SetState(ctx context.Context, svcCtx servicecontext.ServiceContextInterface) op.JobRepository {
	r.ctx = ctx
	r.svcCtx = svcCtx.(*svc.ServiceContext)
	return r
}

func init() {
	app.Register("jobRepository", func() op.JobRepository {
		return NewJobRepository()
	})
}
