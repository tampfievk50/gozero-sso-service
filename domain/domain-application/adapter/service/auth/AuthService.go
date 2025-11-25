package auth

import (
	"context"
	"errors"
	"fmt"
	"gozero-sso-service/domain/domain-application/utils"
	"gozero-sso-service/domain/domain-core/dto"
	"strings"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zeromicro/go-zero/core/logx"
)

func (l *service) RefreshToken(ctx context.Context) (*dto.UserToken, error) {
	//TODO implement me
	panic("implement me")
}

func (l *service) Login(ctx context.Context, userLoginDto *dto.UserLoginDto, config *dto.ConfigDto) (*dto.UserToken, error) {
	user, err := l.userRepo.GetUserByMail(ctx, &userLoginDto.Email)

	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	err = utils.VerifyPassword(user.Password, userLoginDto.Password)

	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	now := time.Now().Unix()
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":      user.ID,
		"username": user.Username,
		"dom":      strings.Trim(strings.Replace(fmt.Sprint(user.ResourceIDs), " ", ",", -1), "[]"),
		"iat":      now,
		"exp":      now + config.Auth.AccessExpire,
		"role":     strings.Trim(strings.Replace(fmt.Sprint(user.RoleIDs), " ", ",", -1), "[]"),
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

func (l *service) HasPermission(ctx context.Context, m *casbin.SyncedEnforcer, sub string, doms []string, path, method string) (bool, error) {
	for _, domainId := range doms {
		enforce, err := m.Enforcer.Enforce(sub, domainId, path, method)
		if err != nil {
			logx.WithContext(ctx).Errorf("casbin error: %v", err.Error())
		}
		if enforce {
			return true, nil
		}
	}
	return false, nil
}
