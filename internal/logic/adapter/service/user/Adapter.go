package user

import (
	"context"
	"github.com/tampfievk50/gozero-core-api/app"
	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"github.com/zeromicro/go-zero/core/logx"
	iport "gozero-sso-service/internal/logic/ports/input/service"
	"gozero-sso-service/internal/svc"
)

type service struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJobService() iport.UserService {
	return &service{}
}

func (l *service) SetState(ctx context.Context, svcCtx servicecontext.ServiceContextInterface) iport.UserService {
	l.ctx = ctx
	l.svcCtx = svcCtx.(*svc.ServiceContext)
	l.Logger = logx.WithContext(ctx)
	return l
}

func init() {
	app.Register("userService", func() iport.UserService {
		return NewJobService()
	})
}
