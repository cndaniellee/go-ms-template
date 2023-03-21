package usercode

import "goms/common/response/errcode"

// 枚举函数
const (
	Register = errcode.User + (iota+10)*100
	Login
	Current
)
