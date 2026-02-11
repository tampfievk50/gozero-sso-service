package dto

type RoleDTO struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ResourceID  *string `json:"resource_id,omitempty"`

	Permissions []PermissionDTO `json:"permissions,omitempty"`
}
