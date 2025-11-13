package dto

type RoleDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ResourceID  *uint  `json:"resource_id,omitempty"`

	Permissions []PermissionDTO `json:"permissions,omitempty"`
}
