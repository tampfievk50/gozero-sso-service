package user

import (
	"gozero-sso-service/internal/logic/ports/output/repository"
	"gozero-sso-service/internal/types"

	"github.com/tampfievk50/gozero-core-api/app"
	"gorm.io/gorm"
)

func (l *service) AddNewJob(jobData *types.AddJob) (*types.Job, error) {
	var job *types.Job
	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		var err error
		jobRepository := app.Make[repository.UserRepository]("jobRepository").
			SetState(l.ctx, l.svcCtx)
		job, err = jobRepository.AddNewJob(tx, jobData)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return job, nil
}
