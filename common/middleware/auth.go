package middleware

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"goms/common/jwtauth"
	"goms/common/response"
	"goms/common/response/errcode"
	"net/http"
)

const (
	AuthRequestUserHeader = "X-User"
)

type AuthRequestUserInfo struct {
	UserId int64 `json:"userId"`
}

/*
获取NginxAuthRequest在Header添加的用户信息
写入Request的Context
*/

type AuthConvertor struct{}

func NewAuthConvertor() *AuthConvertor {
	return &AuthConvertor{}
}

func (c *AuthConvertor) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(AuthRequestUserHeader)
		if header == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// 解析用户信息
		user := &AuthRequestUserInfo{}
		if err := json.Unmarshal([]byte(header), user); err != nil {
			msg := "parse auth info failed"
			logx.Error(errors.Wrap(err, msg))
			response.Write(w, response.ErrResp(0, errcode.Server, response.InternalError, msg), nil)
			return
		}

		// 写入Context
		ctx := r.Context()
		ctx = context.WithValue(ctx, jwtauth.JwtUserIdKey, user.UserId)

		next(w, r.WithContext(ctx))
	}
}
