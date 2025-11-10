package job

import (
	"gozero-sso-service/internal/dataaccess/model"

	"gorm.io/gorm"
)

func (r *repository) GetTaskByName(tx *gorm.DB, taskNames []string) ([]*model.Task, error) {
	var tasks []*model.Task

	if err := tx.Where("name in ?", taskNames).First(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
