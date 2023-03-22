package validator

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type v9 struct {
	validator *validator.Validate
}

func NewV9() httpx.Validator {
	return &v9{validator: validator.New()}
}

func (v v9) Validate(_ *http.Request, data any) error {
	return v.validator.Struct(data)
}
