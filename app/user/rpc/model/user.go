package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
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
	result, err := m.SF.Do(cacheKey, func() (any, error) {
		user := &User{}
		// 读取缓存
		if err := m.Cache.GetCtx(ctx, cacheKey, user); err == nil {
			return user, nil
		}
		// 查询数据
		if err := m.DB.WithContext(ctx).Where("id = ?", id).First(user).Error; err != nil {
			// 无数据写入占位符
			if err == gorm.ErrRecordNotFound {
				if err = m.Cache.SetCtx(ctx, cacheKey, "*"); err != nil {
					logx.WithContext(ctx).Error(errors.Wrap(err, "Cache placeholder failed"))
				}
			}
			return nil, err
		}
		// 写入缓存，缓存错误不影响业务
		if err := m.Cache.SetCtx(ctx, cacheKey, user); err != nil {
			logx.WithContext(ctx).Error(errors.Wrap(err, "Cache data failed"))
		}
		return user, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*User), nil
}

func (m *userModel) Create(ctx context.Context, user *User) error {
	if err := m.DB.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	m.RemoveCache(ctx, user.ID)
	return nil
}
