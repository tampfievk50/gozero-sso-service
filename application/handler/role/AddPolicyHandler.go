package role

import (
	"github.com/jinzhu/copier"
	"gozero-sso-service/application/svc"
	"gozero-sso-service/domain/domain-core/dto"
	"net/http"

	"github.com/tampfievk50/gozero-core-api/servicecontext"

	"gozero-sso-service/application/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddPolicyHandler(svcCtx servicecontext.ServiceContextInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			req               types.RolePermissionRequest
			rolePermissionDto dto.RolePermissionDto
		)
		if err := httpx.Parse(r, &req); err != nil {
			httpx.WriteJson(w, http.StatusBadRequest, types.VResponse(http.StatusBadRequest, err.Error(), nil))
			return
		}

		err := copier.Copy(&rolePermissionDto, &req)
		if err != nil {
			httpx.WriteJson(w, http.StatusBadRequest, types.VResponse(http.StatusBadRequest, err.Error(), nil))
			return
		}
		err = svcCtx.(*svc.ServiceContext).Svc.RoleService.AddPolicy(r.Context(), &rolePermissionDto)
		if err != nil {
			httpx.WriteJson(w, http.StatusInternalServerError, types.VResponse(http.StatusInternalServerError, err.Error(), nil))
		} else {
			httpx.WriteJson(w, http.StatusCreated, types.VResponse(http.StatusCreated, http.StatusText(http.StatusCreated), req))
		}
	}
}
