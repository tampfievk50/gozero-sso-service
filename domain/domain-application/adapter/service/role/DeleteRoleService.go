package role

import (
	"context"
)

func (l *service) DeleteRole(ctx context.Context, id string) error {
	return l.rp.RoleRepository.DeleteRole(ctx, id)
}
