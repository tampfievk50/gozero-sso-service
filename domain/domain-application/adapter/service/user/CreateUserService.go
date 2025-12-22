package user

import (
	"context"
	"errors"
	"gozero-sso-service/domain/domain-core/dto"
	"gozero-sso-service/domain/domain-core/utils"
)

func (l *service) CreateUser(ctx context.Context, userDto *dto.UserDTO) error {
	user, err := l.rp.UserRepository.GetUserByMail(ctx, &userDto.Email)
	if user != nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := utils.HashPassword(userDto.Password)
	if err != nil {
		return err
	}
	userDto.Password = *hashedPassword
	err = l.rp.UserRepository.CreateUser(ctx, userDto)
	if err != nil {
		return err
	}
	return nil
}
