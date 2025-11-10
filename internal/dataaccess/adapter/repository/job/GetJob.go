package job

import (
	"gozero-sso-service/internal/dataaccess/model"
)

func (r *repository) GetJob(id uint) (*model.Job, error) {
	var job model.Job

	err := r.svcCtx.DB.Where("id = ? and is_deleted = ?", id, false).First(&job).Error
	if err != nil {
		return nil, err
	}
	return &job, nil
}
