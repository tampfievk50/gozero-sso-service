package permission

import (
	"gozero-sso-service/application/svc"
	"gozero-sso-service/application/types"
	"gozero-sso-service/domain/domain-core/dto"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateHandler(svcCtx servicecontext.ServiceContextInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			req           types.CreatePermissionRequest
			permissionDto dto.PermissionDTO
		)
		if err := httpx.Parse(r, &req); err != nil {
			httpx.WriteJson(w, http.StatusBadRequest, types.VResponse(http.StatusBadRequest, err.Error(), nil))
			return
		}

		err := copier.Copy(&permissionDto, &req)
		if err != nil {
			httpx.WriteJson(w, http.StatusBadRequest, types.VResponse(http.StatusBadRequest, err.Error(), nil))
			return
		}

		err = svcCtx.(*svc.ServiceContext).Svc.PermissionService.CreatePermission(r.Context(), &permissionDto)
		if err != nil {
			httpx.WriteJson(w, http.StatusInternalServerError, types.VResponse(http.StatusInternalServerError, err.Error(), nil))
		} else {
			httpx.WriteJson(w, http.StatusCreated, types.VResponse(http.StatusCreated, http.StatusText(http.StatusCreated), req))
		}
	}
}
