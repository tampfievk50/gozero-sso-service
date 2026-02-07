package types

type Pager struct {
	Page     int `form:"page,optional"`
	PageSize int `form:"pageSize,optional"`
}
