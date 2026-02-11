package permission

import (
	"context"
	"gozero-sso-service/domain/domain-core/dto"
	"log"
)

const adminRoleName = "Administrator"

func (l *service) CreatePermission(ctx context.Context, permissionDto *dto.PermissionDTO) error {
	if err := l.rp.PermissionRepository.CreatePermission(ctx, permissionDto); err != nil {
		return err
	}

	// Auto-assign new permission to Administrator role as a Casbin policy
	adminRole, err := l.rp.RoleRepository.GetRoleByName(ctx, adminRoleName)
	if err != nil {
		log.Printf("[WARN] Could not find '%s' role for auto-policy: %v", adminRoleName, err)
		return nil
	}

	roleID := adminRole.ID
	// Remove existing policy first to avoid duplicates
	_, _ = l.enforcer.RemoveFilteredPolicy(0, roleID, "0", permissionDto.Path, permissionDto.Action)
	_, err = l.enforcer.AddPolicy(roleID, "0", permissionDto.Path, permissionDto.Action, "allow")
	if err != nil {
		log.Printf("[WARN] Could not auto-assign policy to '%s': %v", adminRoleName, err)
	}

	return nil
}
