package handler

import (
	"gozero-sso-service/internal/handler/account"
	"gozero-sso-service/internal/handler/auth"
	"gozero-sso-service/internal/handler/role"
	"gozero-sso-service/internal/svc"
	"net/http"

	"github.com/tampfievk50/gozero-core-api/servicecontext"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx servicecontext.ServiceContextInterface) {
	server.Use(serverCtx.(svc.ServiceContext).RecoverMiddleware)
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.(svc.ServiceContext).AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/account",
					Handler: account.AccountCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/account",
					Handler: account.AccountGetPageHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/account/:id",
					Handler: account.AccountGetHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/account/:id",
					Handler: account.AccountUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/account/:id",
					Handler: account.AccountDeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/account/:id/state",
					Handler: account.AccountUpdateStateHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.(svc.ServiceContext).Config.Auth.AccessSecret),
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.(svc.ServiceContext).AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/send-email/:id",
					Handler: account.AccountGetHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/test"),
		rest.WithJwt(serverCtx.(svc.ServiceContext).Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/account/email",
				Handler: account.PrivateAccountGetByEmailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/account/ids",
				Handler: account.PrivateAccountGetByIdsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/private/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/auth/forget-password",
				Handler: auth.AuthForgetPasswordHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/auth/is-allow",
				Handler: auth.AuthIsAllowHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/auth/login",
				Handler: auth.AuthLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/auth/login-with-secret",
				Handler: auth.AuthLoginWithSecretHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/auth/resend-otp-with-secret",
				Handler: auth.AuthResendOtpWithSecretHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/auth/reset-password",
				Handler: auth.AuthResetPasswordHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.(svc.ServiceContext).RefreshTokenMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/auth/refresh-token",
					Handler: auth.AuthRefreshTokenHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.(svc.ServiceContext).AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/auth/check-session",
					Handler: auth.AuthCheckSessionHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/auth/logout",
					Handler: auth.AuthLogoutHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/auth/me",
					Handler: auth.AuthGetMeHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.(svc.ServiceContext).Config.Auth.AccessSecret),
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.(svc.ServiceContext).AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/role",
					Handler: role.RoleCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/role",
					Handler: role.RoleGetPageHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/role/:id",
					Handler: role.RoleGetHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/role/:id",
					Handler: role.RoleUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/role/:id",
					Handler: role.RoleDeleteHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.(svc.ServiceContext).Config.Auth.AccessSecret),
		rest.WithPrefix("/v1"),
	)
}
