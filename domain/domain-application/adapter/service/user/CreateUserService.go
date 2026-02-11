package user

import (
	"context"
	"errors"
	"gozero-sso-service/domain/domain-core/dto"
	"gozero-sso-service/domain/domain-core/utils"
	"log"
)

const defaultRoleName = "Authenticated"

func (l *service) CreateUser(ctx context.Context, userDto *dto.UserDTO) error {
	emailExists, err := l.rp.UserRepository.ExistsByEmail(ctx, userDto.Email)
	if err != nil {
		return err
	}
	if emailExists {
		return errors.New("email already exists")
	}

	usernameExists, err := l.rp.UserRepository.ExistsByUsername(ctx, userDto.Username)
	if err != nil {
		return err
	}
	if usernameExists {
		return errors.New("username already exists")
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

	// Auto-assign "Authenticated" role to new user
	defaultRole, err := l.rp.RoleRepository.GetRoleByName(ctx, defaultRoleName)
	if err != nil {
		log.Printf("[WARN] Could not find default role '%s': %v", defaultRoleName, err)
		return nil
	}
	if err = l.rp.UserRepository.AssignRoles(ctx, userDto.ID, []string{defaultRole.ID}, "0"); err != nil {
		log.Printf("[WARN] Could not assign default role to user %s: %v", userDto.ID, err)
	}

	return nil
}
