package response

/**
定义HTTP统一返回MSG，便于前端i18n
客户端无需知悉具体错误信息，只需用错误码帮助定位错误位置
*/

type ErrMsg string

const (
	OK ErrMsg = "OK"

	// 未定义的错误
	UnknownError ErrMsg = "UNKNOWN_ERROR"

	// 参数解析、校验处使用的错误
	InvalidParam ErrMsg = "INVALID_PARAM"
	MissingParam ErrMsg = "MISSING_PARAM"

	// 本服务和依赖服务出现的错误
	InternalError ErrMsg = "INTERNAL_ERROR"
	ServiceError  ErrMsg = "SERVICE_ERROR"

	// 与数据相关的错误
	NoneMatching ErrMsg = "NONE_MATCHING"
	AccessDenied ErrMsg = "ACCESS_DENIED"
)
