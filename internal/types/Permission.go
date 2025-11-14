package types

type CreatePermissionRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Path        string `json:"path" binding:"required"`
	Action      string `json:"action" binding:"required"`
}

type UpdatePermissionRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Path        *string `json:"path,omitempty"`
	Action      *string `json:"action,omitempty"`
}

type PermissionFilterRequest struct {
	Page   int     `form:"page" json:"page"`
	Limit  int     `form:"limit" json:"limit"`
	Search *string `form:"search" json:"search"` // filter by name or path
	Action *string `form:"action" json:"action"`
}
