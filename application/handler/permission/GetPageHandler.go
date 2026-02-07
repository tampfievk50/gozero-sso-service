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

func GetPageHandler(svcCtx servicecontext.ServiceContextInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			req      types.Pager
			pagerDto dto.PagerDto
		)
		if err := httpx.Parse(r, &req); err != nil {
			req = types.Pager{
				Page:     svcCtx.(*svc.ServiceContext).Config.Pager.Page,
				PageSize: svcCtx.(*svc.ServiceContext).Config.Pager.Size,
			}
		}

		_ = copier.Copy(&pagerDto, &req)

		resp, err := svcCtx.(*svc.ServiceContext).Svc.PermissionService.GetAllPermissions(r.Context(), &pagerDto)
		if err != nil {
			httpx.WriteJson(w, http.StatusInternalServerError, types.VResponse(http.StatusInternalServerError, err.Error(), nil))
		} else {
			httpx.WriteJson(w, http.StatusOK, types.VResponse(http.StatusOK, http.StatusText(http.StatusOK), resp))
		}
	}
}
