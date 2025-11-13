package types

type CreateRoleRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	ResourceID  *uint  `json:"resource_id"`
	// If your system supports permissions per role
	PermissionIDs []uint `json:"permission_ids"`
}

type UpdateRoleRequest struct {
	Name          *string `json:"name,omitempty"`
	Description   *string `json:"description,omitempty"`
	ResourceID    *uint   `json:"resource_id,omitempty"`
	PermissionIDs []uint  `json:"permission_ids,omitempty"`
}

type RoleFilterRequest struct {
	Page       int     `form:"page" json:"page"`
	Limit      int     `form:"limit" json:"limit"`
	Search     *string `form:"search" json:"search"` // filter by name
	ResourceID *uint   `form:"resource_id" json:"resource_id"`
}
