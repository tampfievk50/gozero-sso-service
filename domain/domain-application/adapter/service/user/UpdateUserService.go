package user

import (
	"context"
	"gozero-sso-service/domain/domain-core/dto"
)

func (l *service) UpdateUser(ctx context.Context, userDto *dto.UserDTO) error {
	return l.rp.UserRepository.UpdateUser(ctx, userDto)
}
