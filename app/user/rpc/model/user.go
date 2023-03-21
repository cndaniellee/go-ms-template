package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/syncx"
	"goms/common/model"
	"gorm.io/gorm"
	"time"
)

type User struct {
	model.Model
	Username string `gorm:"type:varchar(32);comment:用户名;uniqueIndex"`
	Password string `gorm:"type:varchar(32);comment:密码"`
}

type (
	UserModel interface {
		FindByUsername(ctx context.Context, username string) (*User, error)
		FindById(ctx context.Context, id int64) (*User, error)

		Create(ctx context.Context, user *User) error
	}

	userModel struct {
		db    *gorm.DB
		cache cache.Cache
		table string
	}
)

func NewUserModel(db *gorm.DB, r *redis.Redis) UserModel {
	c := cache.NewNode(r, syncx.NewSingleFlight(), cache.NewStat("user"), gorm.ErrRecordNotFound, cache.WithExpiry(time.Hour), cache.WithNotFoundExpiry(time.Hour))
	return &userModel{db: db, cache: c, table: "user"}
}

func (m *userModel) FindByUsername(ctx context.Context, username string) (*User, error) {
	user := &User{}
	if err := m.db.WithContext(ctx).Where("username = ?", username).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (m *userModel) FindById(ctx context.Context, id int64) (*User, error) {
	user := &User{}
	cacheKey := fmt.Sprintf(model.IdCacheKey, m.table, id)
	// 读取缓存
	if err := m.cache.GetCtx(ctx, cacheKey, user); err == nil {
		return user, nil
	}
	if err := m.db.WithContext(ctx).Where("id = ?", id).First(user).Error; err != nil {
		// 无数据写入占位符
		if err == gorm.ErrRecordNotFound {
			if err = m.cache.SetCtx(ctx, cacheKey, "*"); err != nil {
				logx.WithContext(ctx).Error(errors.Wrap(err, "cache placeholder failed"))
			}
		}
		return nil, err
	}
	// 写入缓存，缓存错误不影响业务
	if err := m.cache.SetCtx(ctx, cacheKey, user); err != nil {
		logx.WithContext(ctx).Error(errors.Wrap(err, "cache data failed"))
	}
	return user, nil
}

func (m *userModel) Create(ctx context.Context, user *User) error {
	if err := m.db.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	// 删除缓存，缓存错误不影响业务
	cacheKey := fmt.Sprintf(model.IdCacheKey, m.table, user.ID)
	if err := m.cache.DelCtx(ctx, cacheKey); err != nil {
		logx.WithContext(ctx).Error(errors.Wrap(err, "cache delete failed"))
	}
	return nil
}
