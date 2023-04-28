package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"goms/common/jwtauth"
	"net/http"
)

/*
使用Redis缓存和检查JwtToken中的ExtraCode
用于单点登录及无效化不再使用的Token
*/

type TokenValidator struct {
	rds *redis.Redis
}

func NewTokenValidator(rds *redis.Redis) *TokenValidator {
	return &TokenValidator{rds: rds}
}

func (m *TokenValidator) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 读取Jwt中的UserId和ExtraCode
		userId, err := r.Context().Value(jwtauth.JwtUserIdKey).(json.Number).Int64()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		extraCode := r.Context().Value(jwtauth.JwtExtraCodeKey)
		if userId == 0 || extraCode == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// 对比Redis中的ExtraCode
		savedExtra, err := m.rds.GetCtx(r.Context(), fmt.Sprintf(jwtauth.CacheExtraUserKey, userId))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if extraCode != savedExtra {
			// 返回406状态，因为401状态客户端会尝试Refresh
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}

		next(w, r)
	}
}
