package types

type UserRoleRequest struct {
	UserID   string   `json:"user_id"`
	RoleID   []string `json:"roles"`
	DomainID string   `json:"domain_id"`
}

type AssignRolesRequest struct {
	RoleIDs  []string `json:"roles"`
	DomainID string   `json:"domain_id"`
}

type RemoveRoleRequest struct {
	RoleID   string `json:"role_id"`
	DomainID string `json:"domain_id"`
}
