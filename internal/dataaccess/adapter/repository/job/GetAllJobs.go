package job

import (
	"gozero-sso-service/internal/dataaccess/model"
)

func (r *repository) GetAllJobs() ([]model.Job, error) {
	var jobs []model.Job

	err := r.svcCtx.DB.Where("status = ?", true).Find(&jobs).Error
	if err != nil {
		return nil, err
	}
	return jobs, nil
}
