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

func PaymentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PaymentReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Write(w, response.ErrResp(-2, ordercode.Payment, response.InvalidParam, err.Error()), nil)
			return
		}

		if err := svcCtx.Validate.Struct(req); err != nil {
			response.Write(w, response.ErrResp(-1, ordercode.Payment, response.MissingParam, err.Error()), nil)
			return
		}

		l := order.NewPaymentLogic(r.Context(), svcCtx)
		resp, err := l.Payment(&req)
		response.Write(w, err, resp)
	}
}
