package register

import (
	"github.com/hibiken/asynq"
	"goms/common/logtool"
	"goms/common/message"
	"goms/service/order/mq/internal/handler"
	"goms/service/order/mq/internal/svc"
)

func RegDq(srvCtx *svc.ServiceContext) error {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     srvCtx.Config.Redis.Host,
			Password: srvCtx.Config.Redis.Pass,
		},
		asynq.Config{
			Logger: logtool.NewAsynqLogger(),
		},
	)

	mux := asynq.NewServeMux()

	mux.Handle(message.DqOrderPaymentTimeout, handler.NewPaymentTimeoutHandler(srvCtx))

	return srv.Start(mux)
}
