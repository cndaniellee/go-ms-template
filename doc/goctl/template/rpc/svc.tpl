package svc

import {{.imports}}

type ServiceContext struct {
	Config config.Config

	SqlDB *gorm.DB
	Redis *redis.Redis

    ExampleModel model.ExampleModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 初始化数据库，使用Gorm
	db := storage.NewSqlDb(c.SqlDB)
	if c.Mode == service.DevMode || c.Mode == service.TestMode {
		db.Debug()
		logx.Must(db.AutoMigrate(&model.Order{}, &model.OrderProduct{}))
	}

	// 初始化缓存
	rds, err := redis.NewRedis(c.Redis.RedisConf)
	logx.Must(err)

	return &ServiceContext{
		Config:c,

		SqlDB: db,
		Redis: rds,

        ExampleModel: model.NewExampleModel(sqlDb, rds),
	}
}
