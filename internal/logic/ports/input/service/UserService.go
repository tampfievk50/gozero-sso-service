package service

import (
	"context"
	"gozero-sso-service/internal/types"

	"github.com/tampfievk50/gozero-core-api/servicecontext"
)

type UserService interface {
	SetState(ctx context.Context, svcCtx servicecontext.ServiceContextInterface) UserService
	GetJob(req *types.IdPathRequest) (*types.Job, error)
	AddNewJob(job *types.AddJob) (*types.Job, error)
	AddNewTask(task *types.AddTask) (*types.Task, error)
	StopJob(req *types.IdPathRequest) error
	DeleteJob(req *types.IdPathRequest) error
	GetAllJobs() ([]types.Job, error)
	ResumeJob(req *types.IdPathRequest) error
	StartJob(req *types.IdPathRequest) error
	UpdateJob(req *types.JobUpdate) (*types.JobUpdate, error)
}
