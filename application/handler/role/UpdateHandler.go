package role

import (
	"gozero-sso-service/application/svc"
	"gozero-sso-service/application/types"
	"gozero-sso-service/domain/domain-core/dto"
	"net/http"
	"strconv"

	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/rest/pathvar"
)

func UpdateHandler(svcCtx servicecontext.ServiceContextInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := pathvar.Vars(r)
		idStr := vars["id"]
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			httpx.WriteJson(w, http.StatusBadRequest, types.VResponse(http.StatusBadRequest, "invalid id", nil))
			return
		}

		var req types.UpdateRoleRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.WriteJson(w, http.StatusBadRequest, types.VResponse(http.StatusBadRequest, err.Error(), nil))
			return
		}

		roleDto := dto.RoleDTO{ID: uint(id)}
		if req.Name != nil {
			roleDto.Name = *req.Name
		}
		if req.Description != nil {
			roleDto.Description = *req.Description
		}

		err = svcCtx.(*svc.ServiceContext).Svc.RoleService.UpdateRole(r.Context(), &roleDto)
		if err != nil {
			httpx.WriteJson(w, http.StatusInternalServerError, types.VResponse(http.StatusInternalServerError, err.Error(), nil))
		} else {
			httpx.WriteJson(w, http.StatusOK, types.VResponse(http.StatusOK, http.StatusText(http.StatusOK), req))
		}
	}
}
