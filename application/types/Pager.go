package types

type Pager struct {
	Page       int         `json:"page,omitempty"`
	PageSize   int         `json:"pageSize,omitempty"`
	TotalPage  int         `json:"totalPage,omitempty"`
	TotalCount int         `json:"totalCount,omitempty"`
	Items      interface{} `json:"items,omitempty"`
}
