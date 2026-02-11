package types

type CreateRoleRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,optional"`
	ResourceID  *string `json:"resource_id,optional"`
	// If your system supports permissions per role
	PermissionIDs []string `json:"permission_ids,optional"`
}

type UpdateRoleRequest struct {
	Name          *string  `json:"name,optional"`
	Description   *string  `json:"description,optional"`
	ResourceID    *string  `json:"resource_id,optional"`
	PermissionIDs []string `json:"permission_ids,optional"`
}

type RoleFilterRequest struct {
	Page       int     `form:"page" json:"page"`
	Limit      int     `form:"limit" json:"limit"`
	Search     *string `form:"search" json:"search"` // filter by name
	ResourceID *string `form:"resource_id" json:"resource_id"`
}
