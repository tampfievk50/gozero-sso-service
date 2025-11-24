package types

import "time"

type CreateUserRequest struct {
	Username    string     `json:"username" binding:"required,min=3"`
	Email       string     `json:"email" binding:"required,email"`
	Password    string     `json:"password" binding:"required,min=6"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	Gender      int32      `json:"gender"` // 1=Male, 2=Female, 3=Other
	ResourceIDs []uint     `json:"resources"`
	IsActive    bool       `json:"is_active"`
	IsSupper    bool       `json:"is_supper"`
	RoleIDs     []uint     `json:"roles"` // optional list of roles to assign
}

type UpdateUserRequest struct {
	Username    *string    `json:"username,omitempty"`
	Email       *string    `json:"email,omitempty"`
	Password    *string    `json:"password,omitempty"` // allow optional password change
	FirstName   *string    `json:"first_name,omitempty"`
	LastName    *string    `json:"last_name,omitempty"`
	DateOfBirth *time.Time `json:"date_of_birth,omitempty"`
	Gender      *int32     `json:"gender,omitempty"`
	ResourceID  *uint      `json:"resource_id,omitempty"`
	IsActive    *bool      `json:"is_active,omitempty"`
	IsSupper    *bool      `json:"is_supper,omitempty"`
	RoleIDs     []uint     `json:"role_ids,omitempty"`
}

type UserFilterRequest struct {
	Page       int     `form:"page" json:"page"`
	Limit      int     `form:"limit" json:"limit"`
	Search     *string `form:"search" json:"search"` // search by username or email
	IsActive   *bool   `form:"is_active" json:"is_active"`
	IsSupper   *bool   `form:"is_supper" json:"is_supper"`
	ResourceID *uint   `form:"resource_id" json:"resource_id"`
}
