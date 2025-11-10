package job

import (
	"github.com/jinzhu/copier"
	"gozero-sso-service/internal/dataaccess/model"
	"gozero-sso-service/internal/types"

	"gorm.io/gorm"
)

func (r *repository) AddNewTask(tx *gorm.DB, taskData *types.AddTask) (*types.Task, error) {
	var taskDto types.Task
	task := model.Task{
		Name:        taskData.Name,
		Description: taskData.Description,
		Status:      taskData.Status,
	}

	if err := tx.Create(&task).Error; err != nil {
		return nil, err
	}
	err := copier.Copy(&taskDto, &task)
	if err != nil {
		return nil, err
	}
	return &taskDto, nil
}
