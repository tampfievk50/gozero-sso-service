package user

import (
	"context"
	"gozero-sso-service/internal/dataaccess/dto"
	"gozero-sso-service/internal/types"
)

func (l *service) GetAllUsers(ctx context.Context) ([]dto.UserDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (l *service) GetUser(ctx context.Context, req *types.IdPathRequest) (*dto.UserDTO, error) {
	//TODO implement me
	panic("implement me")
}
