package permission

import (
	"github.com/casbin/casbin/v2"
	iport "gozero-sso-service/domain/domain-core/port/input/service"
	"gozero-sso-service/domain/domain-core/port/output/repository"
)

type service struct {
	rp       *repository.Repository
	enforcer *casbin.SyncedEnforcer
}

func NewPermissionService(rp *repository.Repository, enforcer *casbin.SyncedEnforcer) iport.PermissionService {
	return &service{
		rp:       rp,
		enforcer: enforcer,
	}
}
