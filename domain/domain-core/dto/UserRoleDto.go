package dto

type UserRoleDto struct {
	UserID   uint   `json:"user_id"`
	RoleID   []uint `json:"roles"`
	DomainID uint   `json:"domain_id"`
}
