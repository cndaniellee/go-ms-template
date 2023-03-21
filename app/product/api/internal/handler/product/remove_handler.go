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

func RemoveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Write(w, response.ErrResp(-2, productcode.Remove, response.InvalidParam, err.Error()), nil)
			return
		}

		if err := svcCtx.Validate.Struct(req); err != nil {
			response.Write(w, response.ErrResp(-1, productcode.Remove, response.MissingParam, err.Error()), nil)
			return
		}

		l := product.NewRemoveLogic(r.Context(), svcCtx)
		err := l.Remove(&req)
		response.Write(w, err, nil)
	}
}
