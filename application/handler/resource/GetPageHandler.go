package resource

import (
	"gozero-sso-service/application/svc"
	"gozero-sso-service/dataaccess/model"
	"net/http"

	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetPageHandler(serverCtx servicecontext.ServiceContextInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := serverCtx.(*svc.ServiceContext)

		var resources []model.Resource
		ctx.DB.Connection.Where("is_deleted = ?", false).Find(&resources)

		httpx.OkJsonCtx(r.Context(), w, map[string]interface{}{
			"code": 200,
			"msg":  "OK",
			"data": resources,
		})
	}
}
