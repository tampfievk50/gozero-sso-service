package middleware

import (
	"context"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

type AuthMiddleware struct {
	logx.Logger
	Enforcer *casbin.SyncedEnforcer
}

func NewAuthMiddleware(Enforcer *casbin.SyncedEnforcer) *AuthMiddleware {
	return &AuthMiddleware{
		Logger:   logx.WithContext(context.Background()),
		Enforcer: Enforcer,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if the user has permission to access this endpoint
		roles := r.Context().Value("roles")
		logx.Infof("roles: %v", roles)
		enforce, err := m.Enforcer.Enforce("1", "*", r.URL.Path, r.Method)
		if err != nil {

		}
		logx.Infof("%v", enforce)
		// Passthrough to next handler if need
		next(w, r)
	}
}
