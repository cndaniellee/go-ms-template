// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	product "goms/service/product/api/internal/handler/product"
	"goms/service/product/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthConvertor},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: product.ListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/detail",
					Handler: product.DetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/edit",
					Handler: product.EditHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/remove",
					Handler: product.RemoveHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/product/v1"),
	)
}
