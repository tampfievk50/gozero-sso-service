package dto

type UserRoleDto struct {
	UserID     uint `json:"user_id"`
	RoleID     uint `json:"role_id"`
	ResourceID uint `json:"resource_id"`
}
