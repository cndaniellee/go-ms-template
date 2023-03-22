package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"goms/common/storage"
)

type Config struct {
	zrpc.RpcServerConf

	SqlDB storage.SqlDbConf
}
