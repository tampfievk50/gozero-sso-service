package user

import (
	"github.com/tampfievk50/gozero-core-api/app"
	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"gozero-sso-service/internal/logic/port/input/service"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gozero-sso-service/internal/types"
)

func UpdateHandler(svcCtx servicecontext.ServiceContextInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		userService := app.Make[service.UserService]("userService")
		err := userService.UpdateUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, req)
		}
	}
}
