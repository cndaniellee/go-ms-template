package model

import (
	"goms/common/model"
)

type OrderProduct struct {
	model.Model
	OrderID   int64 `gorm:"comment:'订单ID'"`
	ProductID int64 `gorm:"comment:'商品ID'"`
	Amount    int64 `gorm:"type:bigint(20);comment:数量"`
	Price     int64 `gorm:"type:bigint(20);comment:价格（单位：分）"`
}
