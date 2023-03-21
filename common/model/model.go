package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        int64          `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

const IdCacheKey = "cache:%s:id:%d"
