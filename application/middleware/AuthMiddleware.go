package middleware

import (
	"context"
	"encoding/json"
	"gozero-sso-service/domain/domain-application/adapter/service"
	"gozero-sso-service/domain/domain-core/dto"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type AuthMiddleware struct {
	logx.Logger
	svc *service.Service
}

func NewAuthMiddleware(svc *service.Service) *AuthMiddleware {
	return &AuthMiddleware{
		Logger: logx.WithContext(context.Background()),
		svc:    svc,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			userRoleDtos []dto.UserRoleDto
		)
		roleRaw := r.Context().Value("roles")
		_ = json.Unmarshal([]byte(roleRaw.(string)), &userRoleDtos)

		logx.Infof("request with role: %v", roleRaw)

		verify, _ := m.svc.AuthService.HasPermission(r.Context(), userRoleDtos, r.URL.Path, r.Method)
		if !verify {
			m.Logger.Errorf("Unauthorized for roles: %s, URL: %s, Path: %s", roleRaw, r.URL.Path, r.Method)
			httpx.WriteJson(w, http.StatusForbidden, nil)
			return
		}
		next(w, r)
	}
}
