package types

import "time"

type CreateUserRequest struct {
	Username    string     `json:"username"`
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	FirstName   string     `json:"first_name,optional"`
	LastName    string     `json:"last_name,optional"`
	DateOfBirth *time.Time `json:"date_of_birth,optional"`
	Gender      int32      `json:"gender,optional"` // 1=Male, 2=Female, 3=Other
	ResourceIDs []string   `json:"resources,optional"`
	IsActive    bool       `json:"is_active,optional"`
	IsSupper    bool       `json:"is_supper,optional"`
	RoleIDs     []string   `json:"roles,optional"` // optional list of roles to assign
}

type UpdateUserRequest struct {
	Id          string     `path:"id,omitempty"`
	FirstName   *string    `json:"first_name,optional"`
	LastName    *string    `json:"last_name,optional"`
	DateOfBirth *time.Time `json:"date_of_birth,optional"`
	Gender      *int32     `json:"gender,optional"`
	IsActive    *bool      `json:"is_active,optional"`
}

type UserFilterRequest struct {
	Page       int     `form:"page" json:"page"`
	Limit      int     `form:"limit" json:"limit"`
	Search     *string `form:"search" json:"search"` // search by username or email
	IsActive   *bool   `form:"is_active" json:"is_active"`
	IsSupper   *bool   `form:"is_supper" json:"is_supper"`
	ResourceID *string `form:"resource_id" json:"resource_id"`
}
