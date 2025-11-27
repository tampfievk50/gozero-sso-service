package types

type Pager struct {
	Page     int `form:"page,omitempty"`
	PageSize int `form:"pageSize,omitempty"`
}
