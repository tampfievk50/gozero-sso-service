package Role

import (
	"context"
	"gozero-sso-service/internal/dataaccess/dto"
	"gozero-sso-service/internal/types"
)

func (l *service) GetAllRoles(ctx context.Context) ([]dto.RoleDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (l *service) GetRole(ctx context.Context, req *types.IdPathRequest) (*dto.RoleDTO, error) {
	//TODO implement me
	panic("implement me")
}
