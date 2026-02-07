package types

type CreatePermissionRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,optional"`
	Path        string `json:"path"`
	Action      string `json:"action"`
}

type UpdatePermissionRequest struct {
	Name        *string `json:"name,optional"`
	Description *string `json:"description,optional"`
	Path        *string `json:"path,optional"`
	Action      *string `json:"action,optional"`
}

type PermissionFilterRequest struct {
	Page   int     `form:"page" json:"page"`
	Limit  int     `form:"limit" json:"limit"`
	Search *string `form:"search" json:"search"` // filter by name or path
	Action *string `form:"action" json:"action"`
}
