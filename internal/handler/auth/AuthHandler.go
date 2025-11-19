package auth

import (
	"fmt"
	"github.com/tampfievk50/gozero-core-api/app"
	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"gozero-sso-service/internal/logic/port/input/service"
	"gozero-sso-service/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AuthHandler(svcCtx servicecontext.ServiceContextInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserLoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		userService, err := app.Make[service.UserService]("userService", app.WithParam(r.Context()), app.WithParam(svcCtx))
		fmt.Println(userService)
		authService, err := app.Make[service.AuthService]("AuthService")
		resp, err := authService.Login(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
