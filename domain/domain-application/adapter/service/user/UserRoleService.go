package user

import (
	"context"
	"fmt"
	"gozero-sso-service/domain/domain-core/dto"
)

func (l *service) GetUserRoles(ctx context.Context, userId uint) ([]dto.RoleDTO, error) {
	return l.rp.UserRepository.GetUserRoles(ctx, userId)
}

func (l *service) AssignRoles(ctx context.Context, userId uint, roleIDs []uint, domainID uint) error {
	return l.rp.UserRepository.AssignRoles(ctx, userId, roleIDs, domainID)
}

func (l *service) RemoveRole(ctx context.Context, userId uint, roleID uint, domainID uint) error {
	return l.rp.UserRepository.RemoveRole(ctx, userId, roleID, domainID)
}

func (l *service) GetUserPermissions(ctx context.Context, userId uint) ([]dto.PermissionDTO, error) {
	// Get user's role assignments (includes ResourceID / domain)
	userRoles, err := l.rp.UserRepository.GetUserRoles(ctx, userId)
	if err != nil {
		return nil, err
	}

	// Build a set of (roleID, domainID) pairs from user's assignments
	type roleDomain struct {
		roleID   string
		domainID string
	}
	assignments := make([]roleDomain, 0, len(userRoles))
	for _, ur := range userRoles {
		var domID uint
		if ur.ResourceID != nil {
			domID = *ur.ResourceID
		}
		assignments = append(assignments, roleDomain{
			roleID:   fmt.Sprintf("%d", ur.ID),
			domainID: fmt.Sprintf("%d", domID),
		})
	}

	// Collect permissions from Casbin policies, scoped by user's role-domain pairs
	permMap := make(map[string]dto.PermissionDTO)

	for _, rd := range assignments {
		policies, err := l.enforcer.GetFilteredPolicy(0, rd.roleID)
		if err != nil {
			continue
		}
		for _, policy := range policies {
			// policy: [roleID, domainID, path, action, effect]
			if len(policy) < 4 {
				continue
			}
			policyDomain := policy[1]
			// Only include policies that match user's domain or are wildcard
			if policyDomain != "*" && policyDomain != rd.domainID {
				continue
			}
			key := policyDomain + "|" + policy[2] + "|" + policy[3]
			permMap[key] = dto.PermissionDTO{
				Path:     policy[2],
				Action:   policy[3],
				DomainID: policyDomain,
			}
		}
	}

	// Check if user is super admin — grant all
	user, err := l.rp.UserRepository.GetUser(ctx, &userId)
	if err == nil && user != nil && user.IsSupper {
		allPolicies, err := l.enforcer.GetPolicy()
		if err == nil {
			for _, policy := range allPolicies {
				if len(policy) >= 4 {
					key := policy[1] + "|" + policy[2] + "|" + policy[3]
					permMap[key] = dto.PermissionDTO{
						Path:     policy[2],
						Action:   policy[3],
						DomainID: policy[1],
					}
				}
			}
		}
	}

	permissions := make([]dto.PermissionDTO, 0, len(permMap))
	for _, perm := range permMap {
		permissions = append(permissions, perm)
	}
	return permissions, nil
}
