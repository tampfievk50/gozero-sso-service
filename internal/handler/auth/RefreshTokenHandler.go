package auth

import (
	"github.com/tampfievk50/gozero-core-api/app"
	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"gozero-sso-service/internal/logic/port/input/service"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RefreshTokenHandler(svcCtx servicecontext.ServiceContextInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authService := app.Make[service.AuthService]("authService")
		resp, err := authService.RefreshToken()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
