package svc

import (
	"goms/app/user/rpc/internal/config"
	"goms/common/storage"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config

	SqlDb *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		SqlDb: storage.NewSqlDb(c.Mode, c.SqlDB),
	}
}
