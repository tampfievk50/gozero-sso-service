package dto

type RolePermissionDto struct {
	RoleID   []string `json:"roles"`
	DomainID string   `json:"domain_id"`
	Path     string `json:"path"`
	Action   string `json:"action"`
	Access   string `json:"access"`
}
