package {{.PkgName}}

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	{{.ImportPackages}}
    "goms/common/response"
    "goms/common/response/errcode/{{.PkgName}}code"
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
            response.Write(w, response.ErrResp(-2, {{.PkgName}}code.{{.Call}}, response.InvalidParam, err.Error()), nil)
			return
		}

        if err := svcCtx.Validate.Struct(req); err != nil {
            response.Write(w, response.ErrResp(-1, {{.PkgName}}code.{{.Call}}, response.MissingParam, err.Error()), nil)
            return
        }

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
        {{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
        {{if .HasResp}}response.Write(w, err, resp){{else}}response.Write(w, err, nil){{end}}
	}
}
