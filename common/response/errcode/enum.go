package errcode

/*
枚举错误码，错误码最终格式为六位数：
   10   00    00
  模块  函数  位置
 */

// 枚举模块
const (
	user int = iota + 10
)

// 枚举函数
const (
	Register = iota + user * 100
	Login
	Current
)
