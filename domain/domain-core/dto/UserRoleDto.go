package dto

type UserRoleDto struct {
	UserID     string `json:"user_id"`
	RoleID     string `json:"role_id"`
	ResourceID string `json:"resource_id"`
}
