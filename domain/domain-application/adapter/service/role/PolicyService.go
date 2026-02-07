package role

import (
	"context"
	"errors"
	"fmt"
)

func (l *service) GetAllPolicies(ctx context.Context, roleID string) ([][]string, error) {
	if roleID != "" {
		return l.enforcer.GetFilteredPolicy(0, roleID)
	}
	return l.enforcer.GetPolicy()
}

func (l *service) GetRolePolicies(ctx context.Context, roleID uint) ([][]string, error) {
	return l.enforcer.GetFilteredPolicy(0, fmt.Sprintf("%d", roleID))
}

func (l *service) RemovePolicy(ctx context.Context, roleID string, domainID string, path string, action string) error {
	ok, err := l.enforcer.RemoveFilteredPolicy(0, roleID, domainID, path, action)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("policy not found")
	}
	return nil
}
