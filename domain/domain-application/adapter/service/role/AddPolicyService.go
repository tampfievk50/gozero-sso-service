package role

import (
	"context"
	"errors"
	"fmt"
	"gozero-sso-service/domain/domain-core/dto"
)

func (l *service) AddPolicy(ctx context.Context, rolePermissionDto *dto.RolePermissionDto) error {
	roleDtos, err := l.rp.RoleRepository.ByIDs(ctx, rolePermissionDto.RoleID)
	if err != nil {
		return err
	}

	if len(roleDtos) < len(rolePermissionDto.RoleID) {
		return errors.New("some roles not found")
	}

	rules := make([][]string, 0)
	for _, v := range rolePermissionDto.RoleID {
		rules = append(rules, []string{
			fmt.Sprintf("%d", v),
			fmt.Sprintf("%d", rolePermissionDto.DomainID),
			rolePermissionDto.Path,
			rolePermissionDto.Action,
			rolePermissionDto.Access,
		})
		_, err = l.enforcer.RemoveFilteredPolicy(0,
			fmt.Sprintf("%d", v),
			fmt.Sprintf("%d", rolePermissionDto.DomainID),
			rolePermissionDto.Path,
			rolePermissionDto.Action)
		if err != nil {
			return err
		}
	}

	_, err = l.enforcer.AddPolicies(rules)
	if err != nil {
		return err
	}

	return nil
}
