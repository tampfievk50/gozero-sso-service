package auth

import (
	"github.com/jinzhu/copier"
	"gozero-sso-service/application/svc"
	"gozero-sso-service/application/types"
	"gozero-sso-service/domain/domain-core/dto"
	"net/http"

	"github.com/tampfievk50/gozero-core-api/servicecontext"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// LoginHandler godoc
//
//	@Security		BearerAuth
//	@Summary		User login
//	@Description	User login
//	@Tags			AuthHandler
//	@Accept			json
//	@Produce		json
//	@Param			X-Profile-ID	header		string	true	"Profile ID"
//	@Success		200				{object}	types.Response
//	@Router			/v1/auth/login [get]
func LoginHandler(svcCtx servicecontext.ServiceContextInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			req          types.UserLoginRequest
			userLoginDto dto.UserLoginDto
			configDto    dto.ConfigDto
		)

		svcCtx := svcCtx.(*svc.ServiceContext)

		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.VResponse(http.StatusInternalServerError, err.Error(), nil))
			return
		}

		err := copier.Copy(&userLoginDto, &req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.VResponse(http.StatusInternalServerError, err.Error(), nil))
			return
		}

		err = copier.Copy(&configDto, &svcCtx.Config)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.VResponse(http.StatusInternalServerError, err.Error(), nil))
			return
		}

		resp, err := svcCtx.Svc.AuthService.Login(r.Context(), &userLoginDto, &configDto)

		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.VResponse(http.StatusInternalServerError, err.Error(), nil))
		} else {
			httpx.OkJsonCtx(r.Context(), w, types.VResponse(http.StatusOK, http.StatusText(http.StatusOK), resp))
		}
		return
	}
}
