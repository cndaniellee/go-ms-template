package product

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goms/app/product/api/internal/logic/product"
	"goms/app/product/api/internal/svc"
	"goms/app/product/api/internal/types"

	"goms/common/response"
	"goms/common/response/errcode/productcode"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Write(w, response.ErrResp(0, productcode.List, response.InvalidParam, err.Error()), nil)
			return
		}

		l := product.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List(&req)
		response.Write(w, err, resp)
	}
}
