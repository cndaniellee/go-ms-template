package model

import (
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {

	var err error
	err = db.AutoMigrate(new(User))
	if err != nil {
		panic(err.Error())
	}

}
