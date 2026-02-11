package resource

import (
	"encoding/json"
	"gozero-sso-service/application/svc"
	"gozero-sso-service/application/types"
	"gozero-sso-service/dataaccess/model"
	"net/http"

	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/rest/pathvar"
)

func UpdateHandler(serverCtx servicecontext.ServiceContextInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := pathvar.Vars(r)
		id := vars["id"]

		var req struct {
			Name        *string `json:"name"`
			Description *string `json:"description"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httpx.WriteJson(w, http.StatusBadRequest, types.VResponse(http.StatusBadRequest, err.Error(), nil))
			return
		}

		ctx := serverCtx.(*svc.ServiceContext)
		var res model.Resource
		if err := ctx.DB.Connection.Where("id = ? AND is_deleted = ?", id, false).First(&res).Error; err != nil {
			httpx.WriteJson(w, http.StatusNotFound, types.VResponse(http.StatusNotFound, "resource not found", nil))
			return
		}

		if req.Name != nil {
			res.Name = *req.Name
		}
		if req.Description != nil {
			res.Description = *req.Description
		}

		if err := ctx.DB.Connection.Save(&res).Error; err != nil {
			httpx.WriteJson(w, http.StatusInternalServerError, types.VResponse(http.StatusInternalServerError, err.Error(), nil))
			return
		}

		httpx.OkJsonCtx(r.Context(), w, map[string]interface{}{
			"code": 200,
			"msg":  "OK",
			"data": res,
		})
	}
}
