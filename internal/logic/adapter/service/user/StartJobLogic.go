package user

import (
	"github.com/jinzhu/copier"
	"github.com/tampfievk50/gozero-core-api/app"
	"gozero-sso-service/internal/logic/ports/output/repository"
	"gozero-sso-service/internal/types"
)

func (l *service) StartJob(req *types.IdPathRequest) error {
	var jobDto types.Job
	jobRepository := app.Make[repository.JobRepository]("jobRepository").
		SetState(l.ctx, l.svcCtx)
	job, err := jobRepository.GetJob(req.Id)
	if err != nil {
		return err
	}
	err = copier.Copy(&jobDto, &job)
	if err != nil {
		return err
	}

	return l.svcCtx.JM.StartJob(*jobDto.ID)
}
