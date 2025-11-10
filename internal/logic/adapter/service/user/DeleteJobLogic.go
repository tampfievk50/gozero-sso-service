package user

import (
	"gozero-sso-service/internal/types"
)

func (l *service) DeleteJob(req *types.IdPathRequest) error {
	err := l.svcCtx.JM.DeleteJob(req.Id)
	if err != nil {
		return err
	}
	return nil
}
