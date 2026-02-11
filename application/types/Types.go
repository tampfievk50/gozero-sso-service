package types

type IdPathRequest struct {
	Id string `path:"id,omitempty,optional"`
}

type AddJob struct {
	Name       string   `json:"name"`
	Expression string   `json:"expression"`
	Status     bool     `json:"status"`
	Tasks      []string `json:"tasks"`
}

type Job struct {
	ID         *string `json:"id,omitempty"`
	Name       string  `json:"name"`
	Expression string  `json:"expression"`
	Status     bool    `json:"status"`
	Tasks      []*Task `json:"tasks"`
	TableName  string  `json:"table"`
}

type JobUpdate struct {
	ID         string `path:"id"`
	Expression string `json:"expression,optional"`
	Status     bool   `json:"status,optional"`
	TableName  string `json:"table,optional"`
}

type AddTask struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Status      bool    `json:"status"`
}

type Task struct {
	ID          *string `json:"id,omitempty"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Status      bool    `json:"status"`
}
