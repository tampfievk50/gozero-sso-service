package auth

import (
	"gozero-sso-service/application/svc"
	"gozero-sso-service/application/types"
	"net/http"

	"github.com/tampfievk50/gozero-core-api/servicecontext"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AuthHandler(svcCtx servicecontext.ServiceContextInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserLoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		resp, err := svcCtx.(*svc.ServiceContext).AuthService.Login(r.Context(), nil)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
