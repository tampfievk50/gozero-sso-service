package model

type JobTask struct {
	JobId  uint `json:"job_id" gorm:"primary_key"`
	TaskId uint `json:"task_id" gorm:"primary_key"`
}
