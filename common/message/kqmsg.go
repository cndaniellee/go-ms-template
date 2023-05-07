package message

type KqConf struct {
	Brokers []string
	Topic   string
}

type (
	OrderProduct struct {
		ID     int64 `json:"id"`
		Amount int64 `json:"amount"`
		Price  int64 `json:"price"`
	}

	KqOrderCreateMsg struct {
		UserID    int64          `json:"userId"`
		Products  []OrderProduct `json:"products"`
		Consignee string         `json:"consignee"`
		Phone     string         `json:"phone"`
		Address   string         `json:"address"`
	}
)
