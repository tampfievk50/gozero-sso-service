package user

import (
	"context"
	"gozero-sso-service/dataaccess/model"
	"gozero-sso-service/domain/domain-core/dto"

	"github.com/jinzhu/copier"
)

func (r *repository) GetAllUsers(ctx context.Context, pager *dto.PagerDto) ([]dto.UserDTO, error) {
	var (
		users    []model.User
		userDtos []dto.UserDTO
	)

	err := r.db.Connection.Model(&model.User{}).
		Scopes(r.db.Paginate(pager.Page, pager.PageSize)).
		Find(&users).Error
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&userDtos, &users)
	if err != nil {
		return nil, err
	}
	return userDtos, nil
}

func (r *repository) GetUser(ctx context.Context, id *uint) (*dto.UserDTO, error) {
	var (
		user    model.User
		userDto dto.UserDTO
	)

	err := r.db.Connection.
		First(&user, *id).Error
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&userDto, &user)
	if err != nil {
		return nil, err
	}
	return &userDto, nil
}

func (r *repository) GetUserByMail(ctx context.Context, email *string) (*dto.UserDTO, error) {
	var (
		user    model.User
		userDto dto.UserDTO
	)
	result := r.db.Connection.Preload("UserRoles").
		Where("email = ?", email).
		Where("is_deleted = ?", false).
		Where("is_active = ?", true).
		First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	err := copier.Copy(&userDto, &user)
	if err != nil {
		return nil, err
	}

	return &userDto, nil
}
