package productcode

import "goms/common/response/errcode"

const (
	List = errcode.Product + (iota+10)*100
	Detail
	Edit
	Remove
)
