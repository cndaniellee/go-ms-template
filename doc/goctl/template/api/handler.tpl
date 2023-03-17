package {{.PkgName}}

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	{{.ImportPackages}}
    "goms/common/response"
    "goms/common/response/errcode"
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
            response.Write(w, response.ErrResp(-2, errcode.{{.Call}}, response.InvalidParam), nil)
			return
		}

        if err := svcCtx.Validate.Struct(req); err != nil {
            response.Write(w, response.ErrResp(-1, errcode.{{.Call}}, response.MissingParam), nil)
            return
        }

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
        {{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
        {{if .HasResp}}response.Write(w, err, resp){{else}}response.Write(w, err){{end}}
	}
}
