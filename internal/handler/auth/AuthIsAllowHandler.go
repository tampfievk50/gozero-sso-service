package auth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gozero-sso-service/internal/logic/auth"
	"gozero-sso-service/internal/svc"
	"gozero-sso-service/internal/types"
)

func AuthIsAllowHandler(svcCtx servicecontext.ServiceContextInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AuthIsAllowRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewAuthIsAllowLogic(r.Context(), svcCtx)
		resp, err := l.AuthIsAllow(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
