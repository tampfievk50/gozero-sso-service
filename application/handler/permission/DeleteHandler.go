package permission

import (
	"gozero-sso-service/application/svc"
	"gozero-sso-service/application/types"
	"net/http"

	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteHandler(svcCtx servicecontext.ServiceContextInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdPathRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.WriteJson(w, http.StatusBadRequest, types.VResponse(http.StatusBadRequest, err.Error(), nil))
			return
		}

		err := svcCtx.(*svc.ServiceContext).Svc.PermissionService.DeletePermission(r.Context(), req.Id)
		if err != nil {
			httpx.WriteJson(w, http.StatusInternalServerError, types.VResponse(http.StatusInternalServerError, err.Error(), nil))
		} else {
			httpx.WriteJson(w, http.StatusOK, types.VResponse(http.StatusOK, http.StatusText(http.StatusOK), nil))
		}
	}
}
