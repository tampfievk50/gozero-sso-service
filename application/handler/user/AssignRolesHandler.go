package user

import (
	"gozero-sso-service/application/svc"
	"gozero-sso-service/application/types"
	"net/http"

	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/rest/pathvar"
)

func AssignRolesHandler(svcCtx servicecontext.ServiceContextInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := pathvar.Vars(r)
		idStr := vars["id"]

		var req types.AssignRolesRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.WriteJson(w, http.StatusBadRequest, types.VResponse(http.StatusBadRequest, err.Error(), nil))
			return
		}

		err := svcCtx.(*svc.ServiceContext).Svc.UserService.AssignRoles(r.Context(), idStr, req.RoleIDs, req.DomainID)
		if err != nil {
			httpx.WriteJson(w, http.StatusInternalServerError, types.VResponse(http.StatusInternalServerError, err.Error(), nil))
		} else {
			httpx.WriteJson(w, http.StatusCreated, types.VResponse(http.StatusCreated, http.StatusText(http.StatusCreated), nil))
		}
	}
}
