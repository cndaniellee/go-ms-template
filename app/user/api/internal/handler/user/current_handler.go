package user

import (
	"net/http"

	"goms/app/user/api/internal/logic/user"
	"goms/app/user/api/internal/svc"
	"goms/common/response"
)

func CurrentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewCurrentLogic(r.Context(), svcCtx)
		resp, err := l.Current()
		response.Write(w, err, resp)
	}
}
