package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"goms/common/model"
	"goms/service/product/rpc/model/enum"
	"gorm.io/gorm"
)

type Product struct {
	model.Model
	Title       string               `gorm:"type:varchar(128);comment:标题"`
	Category    enum.ProductCategory `gorm:"type:tinyint(1);comment:类型"`
	Stock       int64                `gorm:"type:bigint(20);comment:库存"`
	Price       int64                `gorm:"type:bigint(20);comment:价格（单位：分）"`
	Description string               `gorm:"type:mediumtext;comment:详情描述"`
}

type (
	ProductModel interface {
		model.IBaseModel
		//List(ctx context.Context, search string, category enum.ProductCategory, page, pageSize int) ([]*Product, int64, error)
		FindById(ctx context.Context, id int64) (*Product, error)

		Upsert(ctx context.Context, user *Product) error
		Delete(ctx context.Context, id int64) error

		ListByIds(ctx context.Context, ids []int64) ([]*Product, error)
	}

	productModel struct {
		*model.BaseModel
	}
)

func NewProductModel(db *gorm.DB, r *redis.Redis) ProductModel {
	return &productModel{BaseModel: model.NewBaseModel(db, r, "order")}
}

//func (m *productModel) List(ctx context.Context, search string, category enum.ProductCategory, page, pageSize int) ([]*Product, int64, error) {
//	var total int64
//	var products []*Product
//	session := m.DB.WithContext(ctx).Model(&Product{})
//	if search != "" {
//		session.Where("title like ?", "%"+search+"%")
//	}
//	if category != 0 {
//		session.Where("category = ?", category)
//	}
//
//	if err := session.Order("created_at desc").Count(&total).Limit(pageSize).Offset((page - 1) * pageSize).Find(&products).Error; err != nil {
//		return nil, 0, err
//	}
//	return products, total, nil
//}

func (m *productModel) FindById(ctx context.Context, id int64) (*Product, error) {
	cacheKey := fmt.Sprintf(model.IdCacheKey, m.Table, id)
	product := &Product{}
	// 读取缓存或查询数据库
	if err := m.Cache.TakeCtx(ctx, product, cacheKey, func(val any) error {
		return m.DB.WithContext(ctx).Where("id = ?", id).First(val).Error
	}); err != nil {
		return nil, err
	}
	return product, nil
}

func (m *productModel) Upsert(ctx context.Context, product *Product) error {
	if product.ID == 0 {
		if err := m.DB.WithContext(ctx).Create(product).Error; err != nil {
			return err
		}
	} else {
		if err := m.DB.WithContext(ctx).Where("id = ?", product.ID).Updates(product).Error; err != nil {
			return err
		}
	}
	m.RemoveCache(product.ID)
	return nil
}

func (m *productModel) Delete(ctx context.Context, id int64) error {
	if err := m.DB.WithContext(ctx).Delete(&Product{}, "id = ?", id).Error; err != nil {
		return err
	}
	m.RemoveCache(id)
	return nil
}

func (m *productModel) ListByIds(ctx context.Context, ids []int64) ([]*Product, error) {
	var products []*Product
	if err := m.DB.WithContext(ctx).Where("id in ?", ids).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
