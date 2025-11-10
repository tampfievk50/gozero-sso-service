package model

import "time"

type Job struct {
	VModel
	Name       string     `json:"name"`
	Expression string     `json:"expression"`
	Status     bool       `json:"status"`
	LastRun    *time.Time `json:"last_run"`
	NextRun    *time.Time `json:"next_run"`
	Tasks      []*Task    `gorm:"many2many:job_task" json:"tasks"`
	TableName  string     `json:"table_name"`
}
