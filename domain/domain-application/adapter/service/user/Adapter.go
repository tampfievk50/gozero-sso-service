package user

import (
	"github.com/casbin/casbin/v2"
	"gozero-sso-service/dataaccess/adapter/repository"
	iport "gozero-sso-service/domain/domain-core/port/input/service"
)

type service struct {
	rp       *repository.Repository
	enforcer *casbin.SyncedEnforcer
}

func NewUserService(rp *repository.Repository, enforcer *casbin.SyncedEnforcer) iport.UserService {
	return &service{
		rp:       rp,
		enforcer: enforcer,
	}
}
