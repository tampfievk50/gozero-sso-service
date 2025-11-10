package user

import (
	"gozero-sso-service/internal/types"
)

func (l *service) StopJob(req *types.IdPathRequest) error {
	return l.svcCtx.JM.StopJob(req.Id)
}
