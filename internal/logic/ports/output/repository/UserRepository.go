package repository

import (
	"context"
	"gozero-sso-service/internal/dataaccess/model"
	"gozero-sso-service/internal/types"

	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"gorm.io/gorm"
)

type UserRepository interface {
	SetState(ctx context.Context, svcCtx servicecontext.ServiceContextInterface) UserRepository
	GetAllJobs() ([]model.Job, error)
	GetJob(id uint) (*model.Job, error)
	AddNewJob(tx *gorm.DB, job *types.AddJob) (*types.Job, error)
	AddNewTask(tx *gorm.DB, taskData *types.AddTask) (*types.Task, error)
	UpdateJob(tx *gorm.DB, jobUpdate *types.JobUpdate) (*types.Job, error)
}
