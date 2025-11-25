package auth

import (
	iport "gozero-sso-service/domain/domain-core/port/input/service"
	oport "gozero-sso-service/domain/domain-core/port/output/repository"

	"github.com/casbin/casbin/v2"
)

type service struct {
	userRepo oport.UserRepository
	m        *casbin.SyncedEnforcer
}

func NewAuthService(userRepo oport.UserRepository, m *casbin.SyncedEnforcer) iport.AuthService {
	return &service{
		userRepo: userRepo,
		m:        m,
	}
}
