package model

import (
	"goms/common/storage"
)

type User struct {
	storage.Model
	Username string `gorm:"type:varchar(32);comment:用户名;uniqueIndex"`
	Password string `gorm:"type:varchar(32);comment:密码"`
}
