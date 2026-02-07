package role

import (
	"context"
)

func (l *service) DeleteRole(ctx context.Context, id uint) error {
	return l.rp.RoleRepository.DeleteRole(ctx, id)
}
