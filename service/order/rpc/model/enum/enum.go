package enum

type OrderStatus int32

const (
	OrderStatusUnpaid OrderStatus = iota + 1
	OrderStatusProcess
	OrderStatusFinish
	OrderStatusClose
)
