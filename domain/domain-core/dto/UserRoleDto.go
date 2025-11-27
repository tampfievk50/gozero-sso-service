package dto

type UserRoleDto struct {
	UserID     int `json:"user_id"`
	RoleID     int `json:"role_id"`
	ResourceID int `json:"resource_id"`
}
