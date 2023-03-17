package reply

// Reply 实现error用于统一RPC返回的错误
type Reply struct {
	Msg  ErrMsg `json:"msg"`
	Note string `json:"note"`
}

func (r Reply) Error() string {
	return string(r.Msg)
}
