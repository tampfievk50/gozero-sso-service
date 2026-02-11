package permission

import (
	"context"
)

func (l *service) DeletePermission(ctx context.Context, id string) error {
	return l.rp.PermissionRepository.DeletePermission(ctx, id)
}
