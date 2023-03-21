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
		err := sqlDb.AutoMigrate(&model.Example{})
		if err != nil {
			panic(err.Error())
		}
	}

	// 初始化缓存
	red, err := redis.NewRedis(c.Redis.RedisConf)
	if err != nil {
		panic(err.Error())
	}

	return &ServiceContext{
		Config:c,

        ExampleModel: model.NewExampleModel(sqlDb, red),
	}
}
