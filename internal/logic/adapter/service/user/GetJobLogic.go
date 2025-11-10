package user

import (
	"github.com/jinzhu/copier"
	"github.com/tampfievk50/gozero-core-api/app"
	"gozero-sso-service/internal/logic/ports/output/repository"
	"gozero-sso-service/internal/types"
)

func (l *service) GetNum() int {
	return 100
}

func (l *service) GetJob(req *types.IdPathRequest) (*types.Job, error) {
	var jobDto types.Job
	jobRepository := app.Make[repository.JobRepository]("jobRepository").
		SetState(l.ctx, l.svcCtx)
	job, err := jobRepository.GetJob(req.Id)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&jobDto, &job)
	if err != nil {
		return nil, err
	}
	return &jobDto, nil
}
