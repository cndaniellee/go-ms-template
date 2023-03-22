package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"goms/common/model"
	"gorm.io/gorm"
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
	UserID      int64       `gorm:"comment:'用户ID'"`
	RefNo       string      `gorm:"type:varchar(32);comment:参考号;uniqueIndex"`
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
		model.IBaseModel
		List(ctx context.Context, userId int64, status OrderStatus, page, pageSize int) ([]*Order, int64, error)
		FindById(ctx context.Context, userId, id int64) (*Order, error)

		UpdateStatus(ctx context.Context, id int64, status OrderStatus) error
	}

	orderModel struct {
		*model.BaseModel
	}
)

func NewOrderModel(db *gorm.DB, r *redis.Redis) OrderModel {
	return &orderModel{BaseModel: model.NewBaseModel(db, r, "order")}
}

func (m *orderModel) List(ctx context.Context, userId int64, status OrderStatus, page, pageSize int) ([]*Order, int64, error) {
	var total int64
	var orders []*Order
	session := m.DB.WithContext(ctx).Model(&Order{}).Where("user_id = ?", userId).Preload("OrderProducts")
	if status != 0 {
		session.Where("status = ?", status)
	}

	if err := session.Order("created_at desc").Count(&total).Limit(pageSize).Offset((page - 1) * pageSize).Find(&orders).Error; err != nil {
		return nil, 0, err
	}
	return orders, total, nil
}

func (m *orderModel) FindById(ctx context.Context, userId, id int64) (*Order, error) {
	order := &Order{}
	// 查询数据
	if err := m.DB.WithContext(ctx).Where("user_id = ? and id = ?", userId, id).First(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (m *orderModel) UpdateStatus(ctx context.Context, id int64, status OrderStatus) error {
	if err := m.DB.WithContext(ctx).Model(&Order{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}
