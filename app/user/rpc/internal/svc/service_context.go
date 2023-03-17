package svc

import (
	"github.com/zeromicro/go-zero/core/service"
	"goms/app/user/rpc/internal/config"
	"goms/app/user/rpc/model"
	"goms/common/storage"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config

	SqlDb *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlDb := storage.NewSqlDb(c.SqlDB)
	if c.Mode == service.DevMode || c.Mode == service.TestMode {
		sqlDb.Debug()
		model.Migration(sqlDb)
	}

	return &ServiceContext{
		Config: c,

		SqlDb: sqlDb,
	}
}
