package middleware

import (
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"gozero-sso-service/application/core"
	"net/http"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

type AuthMiddleware struct {
	logx.Logger
	Enforcer *casbin.SyncedEnforcer
	svc      *core.Service
}

func NewAuthMiddleware(Enforcer *casbin.SyncedEnforcer, svc *core.Service) *AuthMiddleware {
	return &AuthMiddleware{
		Logger:   logx.WithContext(context.Background()),
		Enforcer: Enforcer,
		svc:      svc,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			roleIds []string
			domains []string
		)
		sub := r.Context().Value("role")
		dom := r.Context().Value("dom")
		if strings.Contains(sub.(string), ",") {
			roleIds = strings.Split(sub.(string), ",")
		} else {
			roleIds = append(roleIds, sub.(string))
		}

		if strings.Contains(dom.(string), ",") {
			domains = strings.Split(dom.(string), ",")
		} else {
			domains = append(domains, dom.(string))
		}

		logx.Infof("roles: %v", sub)
		logx.Infof("domain: %v", dom)

		verify, _ := m.svc.AuthService.HasPermission(r.Context(), m.Enforcer, roleIds, domains, r.URL.Path, r.Method)
		if !verify {
			m.Logger.Errorf("Unauthorized for sub: %s, domain: %s, URL: %s, Path: %s", sub, dom, r.URL.Path, r.Method)
			httpx.WriteJson(w, http.StatusForbidden, nil)
			return
		}
		next(w, r)
	}
}
