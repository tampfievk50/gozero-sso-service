package auth

import (
	"gozero-sso-service/application/svc"
	"net/http"

	"github.com/tampfievk50/gozero-core-api/servicecontext"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RefreshTokenHandler(svcCtx servicecontext.ServiceContextInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := svcCtx.(*svc.ServiceContext).AuthService.RefreshToken(r.Context())
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
