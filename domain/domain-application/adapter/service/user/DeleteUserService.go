package user

import (
	"context"
)

func (l *service) DeleteUser(ctx context.Context, id string) error {
	return l.rp.UserRepository.DeleteUser(ctx, id)
}
