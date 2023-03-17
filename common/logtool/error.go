package logtool

import (
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"goms/common/reply"
)

// CheckRpcConnErr 检查RPC返回的错误，如果是连接错误需要写入日志
func CheckRpcConnErr(logger logx.Logger, err error) (msg reply.ErrMsg, note string) {
	if e, ok := err.(reply.Reply); ok {
		msg = e.Msg
		note = e.Note
	} else {
		logger.Error(errors.Wrapf(err, "user rpc call error"))
	}
	return
}
