package types

type PolicyResponse struct {
	RoleID   string `json:"role_id"`
	DomainID string `json:"domain_id"`
	Path     string `json:"path"`
	Action   string `json:"action"`
	Effect   string `json:"effect"`
}

type DeletePolicyRequest struct {
	RoleID   string `json:"role_id"`
	DomainID string `json:"domain_id"`
	Path     string `json:"path"`
	Action   string `json:"action"`
}
