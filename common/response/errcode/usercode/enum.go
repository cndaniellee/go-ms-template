package usercode

import "goms/common/response/errcode"

// 枚举函数
const (
	Current = errcode.User + (iota+10)*100
)
