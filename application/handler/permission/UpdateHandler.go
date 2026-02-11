package permission

import (
	"gozero-sso-service/application/svc"
	"gozero-sso-service/application/types"
	"gozero-sso-service/domain/domain-core/dto"
	"net/http"

	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/rest/pathvar"
)

func UpdateHandler(svcCtx servicecontext.ServiceContextInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := pathvar.Vars(r)
		idStr := vars["id"]

		var req types.UpdatePermissionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.WriteJson(w, http.StatusBadRequest, types.VResponse(http.StatusBadRequest, err.Error(), nil))
			return
		}

		permissionDto := dto.PermissionDTO{ID: idStr}
		if req.Name != nil {
			permissionDto.Name = *req.Name
		}
		if req.Description != nil {
			permissionDto.Description = *req.Description
		}
		if req.Path != nil {
			permissionDto.Path = *req.Path
		}
		if req.Action != nil {
			permissionDto.Action = *req.Action
		}

		err := svcCtx.(*svc.ServiceContext).Svc.PermissionService.UpdatePermission(r.Context(), &permissionDto)
		if err != nil {
			httpx.WriteJson(w, http.StatusInternalServerError, types.VResponse(http.StatusInternalServerError, err.Error(), nil))
		} else {
			httpx.WriteJson(w, http.StatusOK, types.VResponse(http.StatusOK, http.StatusText(http.StatusOK), req))
		}
	}
}
