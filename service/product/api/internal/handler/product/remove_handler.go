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

func RemoveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Write(w, response.ErrResp(0, productcode.Remove, response.InvalidParam, err.Error()), nil)
			return
		}

		l := product.NewRemoveLogic(r.Context(), svcCtx)
		err := l.Remove(&req)
		response.Write(w, err, nil)
	}
}
