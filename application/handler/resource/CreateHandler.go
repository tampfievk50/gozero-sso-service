package resource

import (
	"gozero-sso-service/application/svc"
	"gozero-sso-service/application/types"
	"gozero-sso-service/dataaccess/model"
	"net/http"

	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateHandler(serverCtx servicecontext.ServiceContextInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Name        string `json:"name"`
			Description string `json:"description"`
		}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.WriteJson(w, http.StatusBadRequest, types.VResponse(http.StatusBadRequest, err.Error(), nil))
			return
		}
		if req.Name == "" {
			httpx.WriteJson(w, http.StatusBadRequest, types.VResponse(http.StatusBadRequest, "name is required", nil))
			return
		}

		ctx := serverCtx.(*svc.ServiceContext)
		resource := model.Resource{Name: req.Name, Description: req.Description}
		if err := ctx.DB.Connection.Create(&resource).Error; err != nil {
			httpx.WriteJson(w, http.StatusInternalServerError, types.VResponse(http.StatusInternalServerError, err.Error(), nil))
			return
		}

		httpx.OkJsonCtx(r.Context(), w, map[string]interface{}{
			"code": 201,
			"msg":  "Created",
			"data": resource,
		})
	}
}
