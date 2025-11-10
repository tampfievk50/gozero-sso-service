package job

import (
	"github.com/jinzhu/copier"
	"gozero-sso-service/internal/dataaccess/model"
	"gozero-sso-service/internal/types"

	"gorm.io/gorm"
)

func (r *repository) AddNewJob(tx *gorm.DB, jobData *types.AddJob) (*types.Job, error) {
	var jobDto types.Job

	tasks, err := r.GetTaskByName(tx, jobData.Tasks)
	if err != nil {
		return nil, err
	}

	job := model.Job{
		Name:       jobData.Name,
		Expression: jobData.Expression,
		Status:     jobData.Status,
		Tasks:      tasks,
	}

	if err := tx.Create(&job).Error; err != nil {
		return nil, err
	}

	err = copier.Copy(&jobDto, &job)
	if err != nil {
		return nil, err
	}
	return &jobDto, nil
}
