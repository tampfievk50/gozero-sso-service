package user

import (
	"github.com/tampfievk50/gozero-core-api/app"
	"gorm.io/gorm"
	"gozero-sso-service/internal/logic/ports/output/repository"
	"gozero-sso-service/internal/types"
)

func (l *service) AddNewTask(task *types.AddTask) (*types.Task, error) {
	var taskDto *types.Task
	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		var err error
		jobRepository := app.Make[repository.JobRepository]("jobRepository").
			SetState(l.ctx, l.svcCtx)
		taskDto, err = jobRepository.AddNewTask(tx, task)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return taskDto, nil
}
