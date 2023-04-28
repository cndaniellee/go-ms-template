package token

import (
	"encoding/json"
	"goms/common/middleware"
	"net/http"

	"goms/service/auth/api/internal/logic/token"
	"goms/service/auth/api/internal/svc"
)

func VerifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := token.NewVerifyLogic(r.Context(), svcCtx)
		// 向Nginx返回状态码和用户信息
		if resp, err := l.Verify(); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
		} else if userData, err := json.Marshal(resp); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Header().Set(middleware.AuthRequestUserHeader, string(userData))
			w.WriteHeader(http.StatusOK)
		}
	}
}
