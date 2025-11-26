package user

import (
	"context"
	"gozero-sso-service/dataaccess/model"
	"gozero-sso-service/domain/domain-core/dto"

	"github.com/jinzhu/copier"
)

func (r *repository) GetAllUsers(ctx context.Context, pager *dto.PagerDto) ([]dto.UserDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) GetUser(ctx context.Context, id *uint) (*dto.UserDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) GetUserByMail(ctx context.Context, email *string) (*dto.UserDTO, error) {
	var (
		user    model.User
		userDto dto.UserDTO
	)
	result := r.db.Connection.Preload("Resources").Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	err := copier.Copy(&userDto, &user)
	if err != nil {
		return nil, err
	}

	for _, id := range user.Resources {
		userDto.ResourceIDs = append(userDto.ResourceIDs, id.ID)
	}

	return &userDto, nil
}
