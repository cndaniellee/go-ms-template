package svc

import (
	"goms/app/user/api/internal/config"
	"gopkg.in/go-playground/validator.v9"
)

type ServiceContext struct {
	Config   config.Config
	Validate *validator.Validate
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Validate: validator.New(),
	}
}
