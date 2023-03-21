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

func EditHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EditReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Write(w, response.ErrResp(-2, productcode.Edit, response.InvalidParam, err.Error()), nil)
			return
		}

		if err := svcCtx.Validate.Struct(req); err != nil {
			response.Write(w, response.ErrResp(-1, productcode.Edit, response.MissingParam, err.Error()), nil)
			return
		}

		l := product.NewEditLogic(r.Context(), svcCtx)
		resp, err := l.Edit(&req)
		response.Write(w, err, resp)
	}
}
