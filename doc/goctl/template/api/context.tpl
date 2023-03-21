package svc

import (
	{{.configImport}}
	"gopkg.in/go-playground/validator.v9"
)

type ServiceContext struct {
	Config {{.config}}
	Validate *validator.Validate
	{{.middleware}}
}

func NewServiceContext(c {{.config}}) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Validate: validator.New(),
		{{.middlewareAssignment}}
	}
}
