package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"goms/app/order/rpc/internal/config"
	"goms/app/order/rpc/model"
	"goms/common/storage"
)

type ServiceContext struct {
	Config config.Config

	OrderModel model.OrderModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 初始化数据库，使用Gorm
	sqlDb := storage.NewSqlDb(c.SqlDB)
	if c.Mode == service.DevMode || c.Mode == service.TestMode {
		sqlDb.Debug()
		logx.Must(sqlDb.AutoMigrate(&model.Order{}, &model.OrderProduct{}))
	}

	// 初始化缓存
	rds, err := redis.NewRedis(c.Redis.RedisConf)
	logx.Must(err)

	return &ServiceContext{
		Config: c,

		OrderModel: model.NewOrderModel(sqlDb, rds),
	}
}
