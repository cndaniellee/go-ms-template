package user

import (
	"net/http"

	"goms/common/response"
	"goms/service/user/api/internal/logic/user"
	"goms/service/user/api/internal/svc"
)

func CurrentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewCurrentLogic(r.Context(), svcCtx)
		resp, err := l.Current()
		response.Write(w, err, resp)
	}
}
