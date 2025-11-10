package auth

import (
	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gozero-sso-service/internal/logic/auth"
	"gozero-sso-service/internal/svc"
)

func AuthRefreshTokenHandler(svcCtx servicecontext.ServiceContextInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := auth.NewAuthRefreshTokenLogic(r.Context(), svcCtx)
		resp, err := l.AuthRefreshToken()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
