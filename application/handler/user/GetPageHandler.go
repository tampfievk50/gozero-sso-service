package user

import (
	"gozero-sso-service/application/svc"
	"net/http"

	"github.com/tampfievk50/gozero-core-api/servicecontext"

	"gozero-sso-service/application/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetPageHandler(svcCtx servicecontext.ServiceContextInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Pager
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		resp, err := svcCtx.(*svc.ServiceContext).UserService.GetAllUsers(r.Context())
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
