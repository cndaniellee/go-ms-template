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

type OrderStatus int32

const (
	OrderStatusUnpaid OrderStatus = iota + 1
	OrderStatusProcess
	OrderStatusFinish
	OrderStatusClose
)

type Order struct {
	model.Model
	Status      OrderStatus `gorm:"type:tinyint(1);comment:状态"`
	Consignee   string      `gorm:"type:varchar(128);comment:收件人"`
	Phone       string      `gorm:"type:varchar(32);comment:手机号"`
	Address     string      `gorm:"type:varchar(128);comment:收件地址"`
	TotalAmount int64       `gorm:"type:bigint(20);comment:总数量"`
	TotalPrice  int64       `gorm:"type:bigint(20);comment:总金额（单位：分）"`

	OrderProducts []*OrderProduct `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;"`
}

type (
	OrderModel interface {
		List(ctx context.Context, status OrderStatus, page, pageSize int) ([]*Order, int64, error)
		FindById(ctx context.Context, id int64) (*Order, error)

		Create(ctx context.Context, user *Order) error
		UpdateStatus(ctx context.Context, id int64, status OrderStatus) error
	}

	orderModel struct {
		db    *gorm.DB
		cache cache.Cache
		sf    syncx.SingleFlight
		table string
	}
)

func NewOrderModel(db *gorm.DB, r *redis.Redis) OrderModel {
	c := cache.NewNode(r, syncx.NewSingleFlight(), cache.NewStat("user"), gorm.ErrRecordNotFound, cache.WithExpiry(time.Hour), cache.WithNotFoundExpiry(time.Hour))
	return &orderModel{db: db, cache: c, sf: syncx.NewSingleFlight(), table: "order"}
}

func (m *orderModel) List(ctx context.Context, status OrderStatus, page, pageSize int) ([]*Order, int64, error) {
	var total int64
	var orders []*Order
	session := m.db.WithContext(ctx).Model(&Order{}).Preload("OrderProducts")
	if status != 0 {
		session.Where("status = ?", status)
	}

	if err := session.Order("created_at desc").Count(&total).Limit(pageSize).Offset((page - 1) * pageSize).Find(&orders).Error; err != nil {
		return nil, 0, err
	}
	return orders, total, nil
}

func (m *orderModel) FindById(ctx context.Context, id int64) (*Order, error) {
	cacheKey := fmt.Sprintf(model.IdCacheKey, m.table, id)
	result, err := m.sf.Do(cacheKey, func() (any, error) {
		order := &Order{}
		// 读取缓存
		if err := m.cache.GetCtx(ctx, cacheKey, order); err == nil {
			return order, nil
		}
		// 查询数据
		if err := m.db.WithContext(ctx).Where("id = ?", id).First(order).Error; err != nil {
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
		if err := m.cache.SetCtx(ctx, cacheKey, order); err != nil {
			logx.WithContext(ctx).Error(errors.Wrap(err, "cache data failed"))
		}
		return order, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*Order), nil
}

func (m *orderModel) Create(ctx context.Context, order *Order) error {
	if err := m.db.WithContext(ctx).Create(order).Error; err != nil {
		return err
	}
	// 删除缓存，缓存错误不影响业务
	cacheKey := fmt.Sprintf(model.IdCacheKey, m.table, order.ID)
	if err := m.cache.DelCtx(ctx, cacheKey); err != nil {
		logx.WithContext(ctx).Error(errors.Wrap(err, "cache delete failed"))
	}
	return nil
}

func (m *orderModel) UpdateStatus(ctx context.Context, id int64, status OrderStatus) error {
	if err := m.db.WithContext(ctx).Model(&Order{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		return err
	}
	// 删除缓存，缓存错误不影响业务
	cacheKey := fmt.Sprintf(model.IdCacheKey, m.table, id)
	if err := m.cache.DelCtx(ctx, cacheKey); err != nil {
		logx.WithContext(ctx).Error(errors.Wrap(err, "cache delete failed"))
	}
	return nil
}
