package role

import (
	"context"

	"gozero-sso-service/dataaccess/model"
	"gozero-sso-service/domain/domain-core/dto"

	"github.com/jinzhu/copier"
)

func (r *repository) CreateRole(ctx context.Context, roleDto *dto.RoleDTO) error {
	var role model.Role

	err := copier.Copy(&role, roleDto)
	if err != nil {
		return err
	}

	return r.db.Connection.Create(&role).Error
}
