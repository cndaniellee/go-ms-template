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

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AuthReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Write(w, response.ErrResp(-2, errcode.Login, response.InvalidParam))
			return
		}

		if err := svcCtx.Validate.Struct(req); err != nil {
			response.Write(w, response.ErrResp(-1, errcode.Login, response.MissingParam))
			return
		}

		l := auth.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		response.Write(w, err, resp)
	}
}
