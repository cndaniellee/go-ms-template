package product

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goms/service/product/api/internal/logic/product"
	"goms/service/product/api/internal/svc"
	"goms/service/product/api/internal/types"

	"goms/common/response"
	"goms/common/response/errcode/productcode"
)

func EditHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EditReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Write(w, response.ErrResp(0, productcode.Edit, response.InvalidParam, err.Error()), nil)
			return
		}

		l := product.NewEditLogic(r.Context(), svcCtx)
		resp, err := l.Edit(&req)
		response.Write(w, err, resp)
	}
}
