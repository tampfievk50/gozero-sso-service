package user

import (
	"context"
	"gozero-sso-service/domain/domain-core/dto"
)

func (l *service) GetAllUsers(ctx context.Context, pager *dto.PagerDto) ([]dto.UserDTO, error) {
	return l.rp.UserRepository.GetAllUsers(ctx, pager)
}

func (l *service) GetUser(ctx context.Context, id uint) (*dto.UserDTO, error) {
	return l.rp.UserRepository.GetUser(ctx, &id)
}
