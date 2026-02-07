package types

type UserRoleRequest struct {
	UserID   uint   `json:"user_id"`
	RoleID   []uint `json:"roles"`
	DomainID uint   `json:"domain_id"`
}

type AssignRolesRequest struct {
	RoleIDs  []uint `json:"roles"`
	DomainID uint   `json:"domain_id"`
}

type RemoveRoleRequest struct {
	RoleID   uint `json:"role_id"`
	DomainID uint `json:"domain_id"`
}
