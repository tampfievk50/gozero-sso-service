package role

import (
	"gozero-sso-service/application/svc"
	"gozero-sso-service/application/types"
	"net/http"

	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeletePolicyHandler(svcCtx servicecontext.ServiceContextInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeletePolicyRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.WriteJson(w, http.StatusBadRequest, types.VResponse(http.StatusBadRequest, err.Error(), nil))
			return
		}

		err := svcCtx.(*svc.ServiceContext).Svc.RoleService.RemovePolicy(r.Context(), req.RoleID, req.DomainID, req.Path, req.Action)
		if err != nil {
			httpx.WriteJson(w, http.StatusInternalServerError, types.VResponse(http.StatusInternalServerError, err.Error(), nil))
		} else {
			httpx.WriteJson(w, http.StatusOK, types.VResponse(http.StatusOK, http.StatusText(http.StatusOK), nil))
		}
	}
}
