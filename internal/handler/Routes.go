package handler

import (
	"gozero-sso-service/internal/handler/auth"
	"gozero-sso-service/internal/handler/role"
	"gozero-sso-service/internal/handler/user"
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
					Handler: user.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/account",
					Handler: user.GetPageHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/account/:id",
					Handler: user.GetHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/account/:id",
					Handler: user.UpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/account/:id",
					Handler: user.DeleteHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.(svc.ServiceContext).Config.Auth.AccessSecret),
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.(svc.ServiceContext).RefreshTokenMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/auth/refresh-token",
					Handler: auth.RefreshTokenHandler(serverCtx),
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
					Method:  http.MethodPost,
					Path:    "/role",
					Handler: role.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/role",
					Handler: role.GetPageHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/role/:id",
					Handler: role.GetHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/role/:id",
					Handler: role.UpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/role/:id",
					Handler: role.DeleteHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.(svc.ServiceContext).Config.Auth.AccessSecret),
		rest.WithPrefix("/v1"),
	)
}
