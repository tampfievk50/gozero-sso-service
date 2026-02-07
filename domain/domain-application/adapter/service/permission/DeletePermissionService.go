package permission

import (
	"context"
)

func (l *service) DeletePermission(ctx context.Context, id uint) error {
	return l.rp.PermissionRepository.DeletePermission(ctx, id)
}
