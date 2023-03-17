package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"goms/common/auth"
)

type Config struct {
	rest.RestConf

	JwtAuth auth.JwtAuth

	UserRpcConf zrpc.RpcClientConf
}
