package svc

import (
	{{.configImport}}
	"gopkg.in/go-playground/validator.v9"
)

type ServiceContext struct {
	Config {{.config}}
	{{.middleware}}
	Validate *validator.Validate
}

func NewServiceContext(c {{.config}}) *ServiceContext {
	return &ServiceContext{
		Config: c,
		{{.middlewareAssignment}}
		Validate: validator.New(),
	}
}
