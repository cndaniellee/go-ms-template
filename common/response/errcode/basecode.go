package errcode

/*
枚举错误码，用于快速定位错误
错误码最终格式为六位数：
   10   00    00
  模块  函数  位置

本文件中枚举模块的数值，**code中枚举函数的数值
*/

const (
	Auth int = (iota + 10) * 10000
	User
	Product
	Order

	Server = 500000
)
