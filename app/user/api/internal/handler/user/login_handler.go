package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goms/app/user/api/internal/logic/user"
	"goms/app/user/api/internal/svc"
	"goms/app/user/api/internal/types"

	"goms/common/response"
	"goms/common/response/errcode/usercode"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AuthReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Write(w, response.ErrResp(0, usercode.Login, response.InvalidParam, err.Error()), nil)
			return
		}

		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		response.Write(w, err, resp)
	}
}
