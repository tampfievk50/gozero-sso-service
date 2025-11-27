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

func UpdateHandler(svcCtx servicecontext.ServiceContextInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			req     types.UpdateUserRequest
			userDto dto.UserDTO
		)
		if err := httpx.Parse(r, &req); err != nil {
			httpx.WriteJson(w, http.StatusInternalServerError, types.VResponse(http.StatusInternalServerError, err.Error(), nil))
			return
		}

		err := copier.Copy(&userDto, &req)
		if err != nil {
			httpx.WriteJson(w, http.StatusInternalServerError, types.VResponse(http.StatusInternalServerError, err.Error(), nil))
			return
		}

		err = svcCtx.(*svc.ServiceContext).Svc.UserService.UpdateUser(r.Context(), &userDto)
		if err != nil {
			httpx.WriteJson(w, http.StatusInternalServerError, types.VResponse(http.StatusInternalServerError, err.Error(), nil))
		} else {
			httpx.WriteJson(w, http.StatusOK, types.VResponse(http.StatusOK, http.StatusText(http.StatusOK), req))
		}
	}
}
