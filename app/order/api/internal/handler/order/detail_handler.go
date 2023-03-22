package order

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goms/app/order/api/internal/logic/order"
	"goms/app/order/api/internal/svc"
	"goms/app/order/api/internal/types"

	"goms/common/response"
	"goms/common/response/errcode/ordercode"
)

func DetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Write(w, response.ErrResp(0, ordercode.Detail, response.InvalidParam, err.Error()), nil)
			return
		}

		l := order.NewDetailLogic(r.Context(), svcCtx)
		resp, err := l.Detail(&req)
		response.Write(w, err, resp)
	}
}
