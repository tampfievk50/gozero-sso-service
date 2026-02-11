package permission

import (
	"context"
	"gozero-sso-service/dataaccess/model"
	"gozero-sso-service/domain/domain-core/dto"

	"github.com/jinzhu/copier"
)

func (r *repository) GetAllPermissions(ctx context.Context, pager *dto.PagerDto) ([]dto.PermissionDTO, error) {
	var (
		permissions    []model.Permission
		permissionDtos []dto.PermissionDTO
	)

	err := r.db.Connection.Model(&model.Permission{}).
		Where("is_deleted = ?", false).
		Scopes(r.db.Paginate(pager.Page, pager.PageSize)).
		Find(&permissions).Error
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&permissionDtos, &permissions)
	if err != nil {
		return nil, err
	}
	return permissionDtos, nil
}

func (r *repository) GetPermission(ctx context.Context, id string) (*dto.PermissionDTO, error) {
	var (
		permission    model.Permission
		permissionDto dto.PermissionDTO
	)

	err := r.db.Connection.First(&permission, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&permissionDto, &permission)
	if err != nil {
		return nil, err
	}
	return &permissionDto, nil
}
