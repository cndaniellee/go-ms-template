package auth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goms/app/user/api/internal/logic/auth"
	"goms/app/user/api/internal/svc"
	"goms/app/user/api/internal/types"

	"goms/common/response"
	"goms/common/response/errcode"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AuthReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Write(w, response.ErrResp(-2, errcode.Register, response.InvalidParam))
			return
		}

		if err := svcCtx.Validate.Struct(req); err != nil {
			response.Write(w, response.ErrResp(-1, errcode.Register, response.MissingParam))
			return
		}

		l := auth.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		response.Write(w, err, resp)
	}
}
