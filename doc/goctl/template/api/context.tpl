package svc

import (
	{{.configImport}}
    "goms/common/validator"
)

type ServiceContext struct {
	Config {{.config}}
	{{.middleware}}
}

func NewServiceContext(c {{.config}}) *ServiceContext {

    // 设置V9校验
    httpx.SetValidator(validator.NewV9())

	return &ServiceContext{
		Config: c,
		{{.middlewareAssignment}}
	}
}
