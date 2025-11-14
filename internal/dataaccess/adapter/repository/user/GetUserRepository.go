package job

import (
	"gozero-sso-service/internal/dataaccess/dto"
	"gozero-sso-service/internal/types"
)

func (r *repository) GetAllUsers(pager *types.Pager) ([]dto.UserDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) GetUser(id *uint) (*dto.UserDTO, error) {
	//TODO implement me
	panic("implement me")
}
