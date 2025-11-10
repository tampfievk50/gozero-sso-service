package user

import (
	"github.com/jinzhu/copier"
	"github.com/tampfievk50/gozero-core-api/app"
	"gozero-sso-service/internal/logic/ports/output/repository"
	"gozero-sso-service/internal/types"
)

func (l *service) GetAllJobs() ([]types.Job, error) {
	var jobDtos []types.Job
	jobRepository := app.Make[repository.JobRepository]("jobRepository").
		SetState(l.ctx, l.svcCtx)
	jobs, err := jobRepository.GetAllJobs()
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&jobDtos, &jobs)
	if err != nil {
		return nil, err
	}
	return jobDtos, nil
}
