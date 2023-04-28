package token

import (
	"net/http"

	"goms/common/response"
	"goms/service/auth/api/internal/logic/token"
	"goms/service/auth/api/internal/svc"
)

func RefreshHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := token.NewRefreshLogic(r.Context(), svcCtx)
		resp, err := l.Refresh()
		response.Write(w, err, resp)
	}
}
