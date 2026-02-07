package role

import (
	"context"

	"gozero-sso-service/dataaccess/model"
	"gozero-sso-service/domain/domain-core/dto"

	"github.com/jinzhu/copier"
)

func (r *repository) GetAllRoles(ctx context.Context, pager *dto.PagerDto) ([]dto.RoleDTO, error) {
	var (
		roles    []model.Role
		roleDtos []dto.RoleDTO
	)

	err := r.db.Connection.Model(&model.Role{}).
		Where("is_deleted = ?", false).
		Scopes(r.db.Paginate(pager.Page, pager.PageSize)).
		Find(&roles).Error
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&roleDtos, &roles)
	if err != nil {
		return nil, err
	}
	return roleDtos, nil
}

func (r *repository) GetRole(ctx context.Context, id uint) (*dto.RoleDTO, error) {
	var (
		role    model.Role
		roleDto dto.RoleDTO
	)

	err := r.db.Connection.First(&role, id).Error
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&roleDto, &role)
	if err != nil {
		return nil, err
	}
	return &roleDto, nil
}

func (r *repository) GetRoleByName(ctx context.Context, name string) (*dto.RoleDTO, error) {
	var (
		role    model.Role
		roleDto dto.RoleDTO
	)

	err := r.db.Connection.Where("name = ? AND is_deleted = ?", name, false).First(&role).Error
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&roleDto, &role)
	if err != nil {
		return nil, err
	}
	return &roleDto, nil
}

func (r *repository) ByIDs(ctx context.Context, ids []uint) ([]dto.RoleDTO, error) {
	var (
		roles    []model.Role
		roleDtos []dto.RoleDTO
	)

	r.db.Connection.Where("id in (?)", ids).Find(&roles)
	err := copier.Copy(&roleDtos, &roles)
	if err != nil {
		return nil, err
	}
	return roleDtos, nil
}
