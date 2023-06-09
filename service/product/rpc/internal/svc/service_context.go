package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"goms/common/storage"
	"goms/service/product/rpc/internal/config"
	"goms/service/product/rpc/model"
	"goms/service/product/rpc/model/es"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config

	SqlDB *gorm.DB
	Redis *redis.Redis

	ProductModel model.ProductModel
	ProductES    *es.ProductES
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 初始化数据库，使用Gorm
	db := storage.NewSqlDb(c.SqlDB)
	if c.Mode == service.DevMode || c.Mode == service.TestMode {
		db.Debug()
		logx.Must(db.AutoMigrate(&model.Product{}))
	}

	// 初始化缓存
	rds, err := redis.NewRedis(c.Redis.RedisConf)
	logx.Must(err)

	return &ServiceContext{
		Config: c,

		SqlDB: db,
		Redis: rds,

		ProductModel: model.NewProductModel(db, rds),
		ProductES:    es.NewProductES(c.ElasticSearch),
	}
}
