package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"goms/common/auth"
	"goms/common/message"
)

type Config struct {
	rest.RestConf

	JwtAuth auth.JwtAuthConf

	OrderRpcConf   zrpc.RpcClientConf
	ProductRpcConf zrpc.RpcClientConf

	OrderCreateConf message.KqConf
}
