package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"goms/common/config"
)

type Config struct {
	rest.RestConf

	JwtAuth config.JwtAuth

	UserRpcConf zrpc.RpcClientConf
}
