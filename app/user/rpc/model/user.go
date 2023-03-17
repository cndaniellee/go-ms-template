package model

import (
	"goms/common/storage"
)

type User struct {
	storage.Model
	Username string
	Password string
}
