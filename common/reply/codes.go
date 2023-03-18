package reply

/**
扩展GRPC-Status中的Code，分离连接类错误，识别自定义错误
*/

const (
	ServiceError = iota + 20 // 当前服务出现异常
	DependsError             // 依赖服务出现异常

	NoneMatching // 数据不存在
	DataConflict // 数据重复
)
