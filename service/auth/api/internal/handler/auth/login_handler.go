package auth

import (
	"goms/common/response/errcode/authcode"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goms/service/auth/api/internal/logic/auth"
	"goms/service/auth/api/internal/svc"
	"goms/service/auth/api/internal/types"

	"goms/common/response"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AuthReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Write(w, response.ErrResp(0, authcode.Login, response.InvalidParam, err.Error()), nil)
			return
		}

		l := auth.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		response.Write(w, err, resp)
	}
}
