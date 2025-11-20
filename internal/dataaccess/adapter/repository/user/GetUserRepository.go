package job

import (
	"context"
	"gozero-sso-service/internal/dataaccess/dto"
	"gozero-sso-service/internal/types"
)

func (r *repository) GetAllUsers(ctx context.Context, pager *types.Pager) ([]dto.UserDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) GetUser(ctx context.Context, id *uint) (*dto.UserDTO, error) {
	//TODO implement me
	panic("implement me")
}
