package account

import (
	"context"
	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gozero-sso-service/internal/logic/account"
	"gozero-sso-service/internal/svc"
	"gozero-sso-service/internal/types"
)

func AccountGetHandler(svcCtx servicecontext.ServiceContextInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdPathRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		svcCtx.Pushers["notify-send-email"].Push(context.Background(), "hello world")

		l := account.NewAccountGetLogic(r.Context(), svcCtx)
		resp, err := l.AccountGet(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
