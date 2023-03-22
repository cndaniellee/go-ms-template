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

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Write(w, response.ErrResp(0, ordercode.List, response.InvalidParam, err.Error()), nil)
			return
		}

		l := order.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List(&req)
		response.Write(w, err, resp)
	}
}
