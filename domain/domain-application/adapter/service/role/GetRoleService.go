package role

import (
	"context"
	"gozero-sso-service/domain/domain-core/dto"
)

func (l *service) GetAllRoles(ctx context.Context, pager *dto.PagerDto) ([]dto.RoleDTO, error) {
	return l.rp.RoleRepository.GetAllRoles(ctx, pager)
}

func (l *service) GetRole(ctx context.Context, id uint) (*dto.RoleDTO, error) {
	return l.rp.RoleRepository.GetRole(ctx, id)
}
