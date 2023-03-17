package reply

/**
定义HTTP统一返回MSG，便于前端i18n
客户端无需知悉具体错误信息，只需用错误码帮助定位错误位置
*/

type ErrMsg string

const (
	NoneMatching ErrMsg = "NONE_MATCHING"

	DependsError ErrMsg = "DEPENDS_ERROR"

	ServiceError ErrMsg = "SERVICE_ERROR"
)
