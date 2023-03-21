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

type ProductCategory int32

const (
	ProductCategoryTool ProductCategory = iota + 1
	ProductCategoryFood
)

type Product struct {
	model.Model
	Title       string          `gorm:"type:varchar(128);comment:标题"`
	Category    ProductCategory `gorm:"type:tinyint(1);comment:类型"`
	Stock       int64           `gorm:"type:bigint(20);comment:库存"`
	Description string          `gorm:"type:mediumtext;comment:详情描述"`
}

type (
	ProductModel interface {
		List(ctx context.Context, search string, category ProductCategory, page, pageSize int) ([]*Product, int64, error)
		FindById(ctx context.Context, id int64) (*Product, error)

		Upsert(ctx context.Context, user *Product) error
		Delete(ctx context.Context, id int64) error
	}

	productModel struct {
		db    *gorm.DB
		cache cache.Cache
		table string
	}
)

func NewProductModel(db *gorm.DB, r *redis.Redis) ProductModel {
	c := cache.NewNode(r, syncx.NewSingleFlight(), cache.NewStat("user"), gorm.ErrRecordNotFound, cache.WithExpiry(time.Hour), cache.WithNotFoundExpiry(time.Hour))
	return &productModel{db: db, cache: c, table: "product"}
}

func (m *productModel) List(ctx context.Context, search string, category ProductCategory, page, pageSize int) ([]*Product, int64, error) {
	var total int64
	var products []*Product
	session := m.db.WithContext(ctx)
	if search != "" {
		session.Where("title like ?", "%"+search+"%")
	}
	if category != 0 {
		session.Where("category = ?", category)
	}

	if err := session.Order("created_at desc").Count(&total).Limit(pageSize).Offset((page - 1) * pageSize).Find(&products).Error; err != nil {
		return nil, 0, err
	}
	return products, total, nil
}

func (m *productModel) FindById(ctx context.Context, id int64) (*Product, error) {
	product := &Product{}
	cacheKey := fmt.Sprintf(model.IdCacheKey, m.table, id)
	// 读取缓存
	if err := m.cache.GetCtx(ctx, cacheKey, product); err == nil {
		return product, nil
	}
	if err := m.db.WithContext(ctx).Where("id = ?", id).First(product).Error; err != nil {
		// 无数据写入占位符
		if err == gorm.ErrRecordNotFound {
			err = m.cache.SetCtx(ctx, cacheKey, "*")
			if err != nil {
				logx.WithContext(ctx).Error(errors.Wrap(err, "cache placeholder failed"))
			}
		}
		return nil, err
	}
	// 写入缓存，缓存错误不影响业务
	if err := m.cache.SetCtx(ctx, cacheKey, product); err != nil {
		logx.WithContext(ctx).Error(errors.Wrap(err, "cache data failed"))
	}
	return product, nil
}

func (m *productModel) Upsert(ctx context.Context, product *Product) error {
	if product.ID == 0 {
		if err := m.db.WithContext(ctx).Create(product).Error; err != nil {
			return err
		}
	} else {
		if err := m.db.WithContext(ctx).Where("id = ?", product.ID).Updates(product).Error; err != nil {
			return err
		}
	}
	// 删除缓存，缓存错误不影响业务
	cacheKey := fmt.Sprintf(model.IdCacheKey, m.table, product.ID)
	if err := m.cache.DelCtx(ctx, cacheKey); err != nil {
		logx.WithContext(ctx).Error(errors.Wrap(err, "cache delete failed"))
	}
	return nil
}

func (m *productModel) Delete(ctx context.Context, id int64) error {
	if err := m.db.WithContext(ctx).Delete(&Product{}, "id = ?", id).Error; err != nil {
		return err
	}
	// 删除缓存，缓存错误不影响业务
	cacheKey := fmt.Sprintf(model.IdCacheKey, m.table, id)
	if err := m.cache.DelCtx(ctx, cacheKey); err != nil {
		logx.WithContext(ctx).Error(errors.Wrap(err, "cache delete failed"))
	}
	return nil
}
