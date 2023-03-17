package storage

import (
	"goms/common/logtool"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type SqlDbConf struct {
	DSN           string
	MaxOpenConns  int
	MaxIdleConns  int
	MaxIdleTime   int
	SlowThreshold int
}

func NewSqlDb(conf SqlDbConf) *gorm.DB {

	// 创建数据库连接
	conn, err := gorm.Open(mysql.Open(conf.DSN), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err.Error())
	}

	// 配置数据库参数
	sqlDb, err := conn.DB()
	if err != nil {
		panic(err.Error())
	}
	sqlDb.SetMaxOpenConns(conf.MaxOpenConns)
	sqlDb.SetMaxIdleConns(conf.MaxIdleConns)
	sqlDb.SetConnMaxIdleTime(time.Duration(conf.MaxIdleTime) * time.Minute)

	// 配置Session
	db := conn.Session(
		&gorm.Session{
			QueryFields: true,
			PrepareStmt: true,
			NewDB:       true,
			Logger:      logtool.NewSqlDbLogger(conf.SlowThreshold),
		},
	)

	return db
}
