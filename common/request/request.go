package request

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"goms/common/auth"
)

func ParseUserId(ctx context.Context) (userId int64, err error) {
	if value, ok := ctx.Value(auth.JwtUserIdKey).(json.Number); ok {
		if userId, err = value.Int64(); err != nil {
			return
		}
	}
	if userId <= 0 {
		err = errors.Errorf("invalid user id: %v", userId)
	}
	return
}
