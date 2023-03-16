package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goms/app/user/internal/logic/user"
	"goms/app/user/internal/svc"
)

func CurrentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewCurrentLogic(r.Context(), svcCtx)
		resp, err := l.Current()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
