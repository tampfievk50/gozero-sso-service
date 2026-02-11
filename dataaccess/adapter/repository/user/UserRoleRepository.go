package user

import (
	"context"
	"gozero-sso-service/dataaccess/model"
	"gozero-sso-service/domain/domain-core/dto"
)

// roleWithResource is a scan target that captures both role fields and the resource_id from user_role
type roleWithResource struct {
	ID          string `gorm:"column:id"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	ResourceID  string `gorm:"column:resource_id"`
}

func (r *repository) GetUserRoles(ctx context.Context, userId string) ([]dto.RoleDTO, error) {
	var rows []roleWithResource

	err := r.db.Connection.
		Table("role").
		Select("role.id, role.name, role.description, user_role.resource_id").
		Joins("JOIN user_role ON user_role.role_id = role.id").
		Where("user_role.user_id = ?", userId).
		Scan(&rows).Error
	if err != nil {
		return nil, err
	}

	roleDtos := make([]dto.RoleDTO, len(rows))
	for i, row := range rows {
		rid := row.ResourceID
		roleDtos[i] = dto.RoleDTO{
			ID:          row.ID,
			Name:        row.Name,
			Description: row.Description,
			ResourceID:  &rid,
		}
	}
	return roleDtos, nil
}

func (r *repository) AssignRoles(ctx context.Context, userId string, roleIDs []string, domainID string) error {
	for _, roleID := range roleIDs {
		userRole := model.UserRole{
			UserID:     userId,
			RoleID:     roleID,
			ResourceID: domainID,
		}
		// Use FirstOrCreate to avoid duplicates
		err := r.db.Connection.
			Where("user_id = ? AND role_id = ? AND resource_id = ?", userId, roleID, domainID).
			FirstOrCreate(&userRole).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *repository) RemoveRole(ctx context.Context, userId string, roleID string, domainID string) error {
	return r.db.Connection.
		Where("user_id = ? AND role_id = ? AND resource_id = ?", userId, roleID, domainID).
		Delete(&model.UserRole{}).Error
}
