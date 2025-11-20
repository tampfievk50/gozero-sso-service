package user

import (
	"context"
	"gozero-sso-service/domain/domain-core/dto"
)

func (r *repository) GetAllUsers(ctx context.Context, pager *dto.PagerDto) ([]dto.UserDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) GetUser(ctx context.Context, id *uint) (*dto.UserDTO, error) {
	//TODO implement me
	panic("implement me")
}
