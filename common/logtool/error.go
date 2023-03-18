package logtool

import (
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

// CheckRpcConnErr 检查RPC返回的错误，如果是连接错误需要写入日志
func CheckRpcConnErr(logger logx.Logger, err error) (code codes.Code, msg string) {
	log.Println(err.Error())
	if e, ok := status.FromError(err); ok {
		if e.Code() < 20 {
			logger.Error(errors.Wrapf(err, "rpc connection error"))
		} else {
			code = e.Code()
			msg = e.Message()
		}
	} else {
		logger.Error(errors.Wrapf(err, "unknows rpc call error"))
	}
	return
}
