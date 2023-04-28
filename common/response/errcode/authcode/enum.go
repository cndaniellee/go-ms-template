package authcode

import "goms/common/response/errcode"

// 枚举函数
const (
	Register = errcode.Auth + (iota+10)*100
	Login
	Refresh
)
