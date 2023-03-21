package ordercode

import "goms/common/response/errcode"

const (
	Submit = errcode.Order + (iota+10)*100
	List
	Detail
	Payment
)
