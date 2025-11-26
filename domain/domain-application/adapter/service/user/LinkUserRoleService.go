package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/casbin/casbin/v2"
	"gozero-sso-service/domain/domain-core/dto"
)

func (l *service) LinkUserRole(ctx context.Context, userRoleDto *dto.UserRoleDto, m *casbin.SyncedEnforcer) error {
	roleDtos, err := l.rp.RoleRepository.GetRoleByIds(ctx, userRoleDto.RoleID)
	if err != nil {
		return err
	}

	if len(roleDtos) < len(userRoleDto.RoleID) {
		return errors.New("some roles not found")
	}

	rules := make([][]string, 0)
	for _, v := range userRoleDto.RoleID {
		rules = append(rules, []string{
			fmt.Sprintf("%d", userRoleDto.UserID),
			fmt.Sprintf("%d", v),
			fmt.Sprintf("%d", userRoleDto.DomainID),
		})
	}
	_, err = m.AddGroupingPolicies(rules)
	if err != nil {
		return err
	}

	return nil
}
