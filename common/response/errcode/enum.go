package errcode

/*
枚举错误码，错误码最终格式为六位数：
   10   00    00
  模块  函数  位置
用于快速定位错误
 */

// 枚举模块
const (
	user int = (iota + 10) * 10000
)

// 枚举函数
const (
	Register = user + (iota + 10) * 100
	Login
	Current
)
