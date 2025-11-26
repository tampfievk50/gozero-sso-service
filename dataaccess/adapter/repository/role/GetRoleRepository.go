package role

import (
	"context"
	"github.com/jinzhu/copier"
	"gozero-sso-service/dataaccess/model"
	"gozero-sso-service/domain/domain-core/dto"
)

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
