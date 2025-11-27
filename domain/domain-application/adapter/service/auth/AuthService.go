package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gozero-sso-service/domain/domain-application/utils"
	"gozero-sso-service/domain/domain-core/dto"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zeromicro/go-zero/core/logx"
)

func (l *service) RefreshToken(ctx context.Context) (*dto.UserToken, error) {
	//TODO implement me
	panic("implement me")
}

func (l *service) Login(ctx context.Context, userLoginDto *dto.UserLoginDto, config *dto.ConfigDto) (*dto.UserToken, error) {
	role := make([]byte, 0)
	user, err := l.rp.UserRepository.GetUserByMail(ctx, &userLoginDto.Email)

	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	err = utils.VerifyPassword(user.Password, userLoginDto.Password)

	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	if len(user.UserRoles) > 0 {
		role, err = json.Marshal(user.UserRoles)
	}

	now := time.Now().Unix()
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":      user.ID,
		"username": user.Username,
		"isSuper":  user.IsSupper,
		"roles":    string(role),
		"iat":      now,
		"exp":      now + config.Auth.AccessExpire,
	}).SignedString([]byte(config.Auth.AccessSecret))

	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":      user.ID,
		"username": user.Username,
		"iat":      now,
		"exp":      now + config.Auth.RefreshExpire,
	}).SignedString([]byte(config.Auth.RefreshSecret))

	if err != nil {
		return nil, err
	}

	return &dto.UserToken{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (l *service) HasPermission(ctx context.Context, isSuperAdmin bool, userRoleDtos []dto.UserRoleDto, path, method string) (bool, error) {
	if isSuperAdmin {
		enforce, err := l.enforcer.Enforcer.Enforce(
			"-1",
			"*",
			path,
			method)
		if err != nil {
			logx.WithContext(ctx).Errorf("casbin error: %v", err.Error())
		}
		if enforce {
			return true, nil
		}
	} else {
		for _, userRoleDto := range userRoleDtos {
			enforce, err := l.enforcer.Enforcer.Enforce(
				fmt.Sprintf("%v", userRoleDto.RoleID),
				fmt.Sprintf("%v", userRoleDto.ResourceID),
				path,
				method)
			if err != nil {
				logx.WithContext(ctx).Errorf("casbin error: %v", err.Error())
			}
			if enforce {
				return true, nil
			}
		}
	}

	return false, nil
}
