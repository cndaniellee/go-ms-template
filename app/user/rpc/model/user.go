package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"goms/common/model"
	"gorm.io/gorm"
)

type User struct {
	model.Model
	Username string `gorm:"type:varchar(32);comment:用户名;uniqueIndex"`
	Password string `gorm:"type:varchar(32);comment:密码"`
}

type (
	UserModel interface {
		model.IBaseModel
		FindByUsername(ctx context.Context, username string) (*User, error)
		FindById(ctx context.Context, id int64) (*User, error)

		Create(ctx context.Context, user *User) error
	}

	userModel struct {
		*model.BaseModel
	}
)

func NewUserModel(db *gorm.DB, r *redis.Redis) UserModel {
	return &userModel{BaseModel: model.NewBaseModel(db, r, "user")}
}

func (m *userModel) FindByUsername(ctx context.Context, username string) (*User, error) {
	user := &User{}
	if err := m.DB.WithContext(ctx).Where("username = ?", username).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (m *userModel) FindById(ctx context.Context, id int64) (*User, error) {
	cacheKey := fmt.Sprintf(model.IdCacheKey, m.Table, id)
	user := &User{}
	// 读取缓存或查询数据库
	if err := m.Cache.TakeCtx(ctx, user, cacheKey, func(val any) error {
		return m.DB.WithContext(ctx).Where("id = ?", id).First(val).Error
	}); err != nil {
		return nil, err
	}
	return user, nil
}

func (m *userModel) Create(ctx context.Context, user *User) error {
	if err := m.DB.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	m.RemoveCache(user.ID)
	return nil
}
