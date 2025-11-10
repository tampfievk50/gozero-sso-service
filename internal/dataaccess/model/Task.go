package model

import "time"

type Task struct {
	VModel
	Name        string     `json:"name"`
	Description *string    `json:"expression"`
	Status      bool       `json:"status"`
	Version     uint       `json:"version"`
	LastRun     *time.Time `json:"last_run"`
	LastId      uint       `json:"last_id"`
	Jobs        []*Job     `gorm:"many2many:job_task" json:"jobs"`
}
