package types

type RolePermissionRequest struct {
	RoleID   []uint `json:"roles"`
	DomainID uint   `json:"domain_id"`
	Path     string `json:"path"`
	Action   string `json:"action"`
	Access   string `json:"access" binding:"required"`
}
