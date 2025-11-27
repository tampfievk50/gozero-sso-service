package dto

import "time"

type UserDTO struct {
	ID          uint       `json:"id"`
	Username    string     `json:"username"`
	Email       string     `json:"email"`
	Password    string     `json:"-"`
	Avatar      string     `json:"avatar"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	Age         *int32     `json:"age"`
	Gender      int32      `json:"gender"` // 1=Male, 2=Female, 3=Other
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	IsActive    bool       `json:"is_active"`
	IsSupper    bool       `json:"is_supper"`
	LastLogin   *time.Time `json:"last_login"`
	LastIP      string     `json:"last_ip"`

	UserRoles   []UserRoleDto `json:"user_roles,omitempty"`
	RoleIDs     []uint        `json:"roles,omitempty"`
	ResourceIDs []uint        `json:"resources,omitempty"`
}
