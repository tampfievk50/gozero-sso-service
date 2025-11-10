package user

import (
	"github.com/tampfievk50/gozero-core-api/app"
	"gorm.io/gorm"
	"gozero-sso-service/internal/logic/ports/output/repository"
	"gozero-sso-service/internal/types"
)

func (l *service) UpdateJob(jobUpdate *types.JobUpdate) (*types.JobUpdate, error) {
	var job *types.Job
	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		var err error
		jobRepository := app.Make[repository.JobRepository]("jobRepository").
			SetState(l.ctx, l.svcCtx)
		job, err = jobRepository.UpdateJob(tx, jobUpdate)
		if err != nil {
			return err
		}

		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if job != nil {
		if job.Status {
			err = l.svcCtx.JM.StartJob(*job.ID)
		} else {
			err = l.svcCtx.JM.StopJob(*job.ID)
		}
		if err != nil {
			return nil, err
		}
	}
	return jobUpdate, nil
}
