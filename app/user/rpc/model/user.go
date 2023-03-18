package model

import (
	"goms/common/storage"
)

type User struct {
	storage.Model
	Username string `gorm:"unique_index;type:varchar(32);comment:'用户名'"`
	Password string `gorm:"type:varchar(32);comment:'密码'"`
}
