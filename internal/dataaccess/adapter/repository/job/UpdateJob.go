package job

import (
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gozero-sso-service/internal/dataaccess/model"
	"gozero-sso-service/internal/types"
)

func (r *repository) UpdateJob(tx *gorm.DB, jobUpdate *types.JobUpdate) (*types.Job, error) {
	var job model.Job
	var jobDto types.Job

	err := r.svcCtx.DB.Where("id = ?", jobUpdate.ID).Find(&job).Error
	if err != nil {
		return nil, err
	}

	job.Expression = jobUpdate.Expression
	job.Status = jobUpdate.Status
	job.TableName = jobUpdate.TableName

	err = tx.Save(&job).Error
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&jobDto, &job)
	if err != nil {
		return nil, err
	}
	return &jobDto, nil
}
