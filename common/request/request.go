package request

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"goms/common/jwtauth"
)

func ParseUserId(ctx context.Context) (userId int64, err error) {
	if userId, err = ctx.Value(jwtauth.JwtUserIdKey).(json.Number).Int64(); err != nil {
		return
	}
	if userId <= 0 {
		err = errors.Errorf("invalid user id: %v", userId)
	}
	return
}
