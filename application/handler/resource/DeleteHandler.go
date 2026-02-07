package resource

import (
	"gozero-sso-service/application/svc"
	"gozero-sso-service/application/types"
	"gozero-sso-service/dataaccess/model"
	"net/http"
	"strconv"

	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/rest/pathvar"
)

func DeleteHandler(serverCtx servicecontext.ServiceContextInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := pathvar.Vars(r)
		id, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			httpx.WriteJson(w, http.StatusBadRequest, types.VResponse(http.StatusBadRequest, "invalid id", nil))
			return
		}

		ctx := serverCtx.(*svc.ServiceContext)
		result := ctx.DB.Connection.
			Model(&model.Resource{}).
			Where("id = ? AND is_deleted = ?", id, false).
			Update("is_deleted", true)

		if result.Error != nil {
			httpx.WriteJson(w, http.StatusInternalServerError, types.VResponse(http.StatusInternalServerError, result.Error.Error(), nil))
			return
		}
		if result.RowsAffected == 0 {
			httpx.WriteJson(w, http.StatusNotFound, types.VResponse(http.StatusNotFound, "resource not found", nil))
			return
		}

		httpx.OkJsonCtx(r.Context(), w, map[string]interface{}{
			"code": 200,
			"msg":  "OK",
		})
	}
}
