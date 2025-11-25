package middleware

import (
	"context"
	"fmt"
	"gozero-sso-service/domain/domain-application/adapter/service"
	"net/http"
	"strings"

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
			domains []string
		)
		sub := r.Context().Value("uid")
		dom := r.Context().Value("dom")

		if strings.Contains(dom.(string), ",") {
			domains = strings.Split(dom.(string), ",")
		} else {
			domains = append(domains, dom.(string))
		}

		logx.Infof("sub: %v, domain: %v", sub, dom)

		verify, _ := m.svc.AuthService.HasPermission(r.Context(), fmt.Sprintf("%v", sub), domains, r.URL.Path, r.Method)
		if !verify {
			m.Logger.Errorf("Unauthorized for sub: %s, domain: %s, URL: %s, Path: %s", sub, dom, r.URL.Path, r.Method)
			httpx.WriteJson(w, http.StatusForbidden, nil)
			return
		}
		next(w, r)
	}
}
