package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/syncx"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        int64          `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

const IdCacheKey = "Cache:%s:id:%d"
const IdLockKey = "Lock:%s:id:%d"

type (
	IBaseModel interface {
		Name() string
		RemoveCache(ctx context.Context, id int64)
	}

	BaseModel struct {
		DB    *gorm.DB
		Cache cache.Cache
		SF    syncx.SingleFlight
		Table string
	}
)

func NewBaseModel(db *gorm.DB, r *redis.Redis, table string) *BaseModel {
	c := cache.NewNode(r, syncx.NewSingleFlight(), cache.NewStat(table), gorm.ErrRecordNotFound, cache.WithExpiry(time.Hour), cache.WithNotFoundExpiry(time.Hour))
	return &BaseModel{DB: db, Cache: c, SF: syncx.NewSingleFlight(), Table: table}
}

func (m *BaseModel) Name() string {
	return m.Table
}

func (m *BaseModel) RemoveCache(ctx context.Context, id int64) {
	// 删除缓存，缓存错误不影响业务
	cacheKey := fmt.Sprintf(IdCacheKey, m.Table, id)
	if err := m.Cache.DelCtx(ctx, cacheKey); err != nil {
		logx.WithContext(ctx).Error(errors.Wrap(err, "Cache delete failed"))
	}
}
