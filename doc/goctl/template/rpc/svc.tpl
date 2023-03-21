package svc

import {{.imports}}

type ServiceContext struct {
	Config config.Config

    ExampleModel model.ExampleModel
}

func NewServiceContext(c config.Config) *ServiceContext {

    // 初始化数据库，使用Gorm
	sqlDb := storage.NewSqlDb(c.SqlDB)
	if c.Mode == service.DevMode || c.Mode == service.TestMode {
		sqlDb.Debug()
		logx.Must(sqlDb.AutoMigrate(&model.Example{}))
	}

	// 初始化缓存
	rds, err := redis.NewRedis(c.Redis.RedisConf)
	logx.Must(err)

	return &ServiceContext{
		Config:c,

        ExampleModel: model.NewExampleModel(sqlDb, rds),
	}
}
