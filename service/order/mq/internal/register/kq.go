package register

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"goms/service/order/mq/internal/handler"
	"goms/service/order/mq/internal/svc"
)

func RegKq(group *service.ServiceGroup, svcCtx *svc.ServiceContext) error {

	// 注册Kafka消费者
	orderCreateJob, err := kq.NewQueue(svcCtx.Config.OrderCreateConf, handler.NewOrderCreateHandler(svcCtx))
	if err != nil {
		return err
	}
	group.Add(orderCreateJob)

	return nil
}
