package user

import (
	"github.com/jinzhu/copier"
	"gozero-sso-service/application/svc"
	"gozero-sso-service/domain/domain-core/dto"
	"net/http"

	"github.com/tampfievk50/gozero-core-api/servicecontext"

	"gozero-sso-service/application/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LinkUserRoleHandler(svcCtx servicecontext.ServiceContextInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			req         types.UserRoleRequest
			userRoleDto dto.UserRoleDto
		)
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, err)
			return
		}

		err := copier.Copy(&userRoleDto, &req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, err)
			return
		}
		svcCtx := svcCtx.(*svc.ServiceContext)
		err = svcCtx.Svc.UserService.LinkUserRole(r.Context(), &userRoleDto)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.VResponse(http.StatusInternalServerError, err.Error(), nil))
		} else {
			httpx.OkJsonCtx(r.Context(), w, types.VResponse(http.StatusCreated, http.StatusText(http.StatusCreated), req))
		}
	}
}
