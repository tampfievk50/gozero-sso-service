package role

import (
	"gozero-sso-service/application/svc"
	"gozero-sso-service/application/types"
	"net/http"

	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetPoliciesHandler(svcCtx servicecontext.ServiceContextInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roleID := r.URL.Query().Get("role_id")

		policies, err := svcCtx.(*svc.ServiceContext).Svc.RoleService.GetAllPolicies(r.Context(), roleID)
		if err != nil {
			httpx.WriteJson(w, http.StatusInternalServerError, types.VResponse(http.StatusInternalServerError, err.Error(), nil))
			return
		}

		var resp []types.PolicyResponse
		for _, p := range policies {
			pr := types.PolicyResponse{
				RoleID: p[0],
			}
			if len(p) > 1 {
				pr.DomainID = p[1]
			}
			if len(p) > 2 {
				pr.Path = p[2]
			}
			if len(p) > 3 {
				pr.Action = p[3]
			}
			if len(p) > 4 {
				pr.Effect = p[4]
			}
			resp = append(resp, pr)
		}

		httpx.WriteJson(w, http.StatusOK, types.VResponse(http.StatusOK, http.StatusText(http.StatusOK), resp))
	}
}

func GetRolePoliciesHandler(svcCtx servicecontext.ServiceContextInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdPathRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.WriteJson(w, http.StatusBadRequest, types.VResponse(http.StatusBadRequest, err.Error(), nil))
			return
		}

		policies, err := svcCtx.(*svc.ServiceContext).Svc.RoleService.GetRolePolicies(r.Context(), req.Id)
		if err != nil {
			httpx.WriteJson(w, http.StatusInternalServerError, types.VResponse(http.StatusInternalServerError, err.Error(), nil))
			return
		}

		var resp []types.PolicyResponse
		for _, p := range policies {
			pr := types.PolicyResponse{
				RoleID: p[0],
			}
			if len(p) > 1 {
				pr.DomainID = p[1]
			}
			if len(p) > 2 {
				pr.Path = p[2]
			}
			if len(p) > 3 {
				pr.Action = p[3]
			}
			if len(p) > 4 {
				pr.Effect = p[4]
			}
			resp = append(resp, pr)
		}

		httpx.WriteJson(w, http.StatusOK, types.VResponse(http.StatusOK, http.StatusText(http.StatusOK), resp))
	}
}
