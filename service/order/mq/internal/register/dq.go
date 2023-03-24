package register

import (
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"goms/service/order/mq/internal/handler/dq"
	"goms/service/order/mq/internal/svc"
)

func RegDq(srvCtx *svc.ServiceContext) error {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     srvCtx.Config.Redis.Host,
			Password: srvCtx.Config.Redis.Pass,
		},
		asynq.Config{
			IsFailure: func(err error) bool {
				if err != nil {
					logx.Error(errors.Wrap(err, "asynq handler error"))
					return true
				}
				return false
			},
		},
	)

	mux := asynq.NewServeMux()

	mux.Handle("", dq.NewPaymentTimeoutHandler(srvCtx))

	return srv.Run(mux)
}
