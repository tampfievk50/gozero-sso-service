package role

import (
	"context"
	"gozero-sso-service/domain/domain-core/dto"
)

func (l *service) CreateRole(ctx context.Context, roleDto *dto.RoleDTO) error {
	return l.rp.RoleRepository.CreateRole(ctx, roleDto)
}
