package message

const DqOrderPaymentTimeout = "dq:order:timeout:payment"

type DqOrderIdMsg struct {
	OrderID int64 `json:"orderId"`
}
