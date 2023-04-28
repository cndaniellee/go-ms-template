package svc

import (
	{{.configImport}}
)

type ServiceContext struct {
	Config {{.config}}

	AuthConvertor rest.Middleware
}

func NewServiceContext(c {{.config}}) *ServiceContext {

	return &ServiceContext{
		Config: c,

        AuthConvertor: middleware.NewAuthConvertor().Handle,
	}
}
