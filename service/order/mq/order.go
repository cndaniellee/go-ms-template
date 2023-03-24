package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"goms/service/order/mq/internal/config"
	"goms/service/order/mq/internal/register"
	"goms/service/order/mq/internal/svc"
)

var configFile = flag.String("f", "etc/order.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	logx.Must(c.SetUp())

	group := service.NewServiceGroup()
	defer group.Stop()

	ctx := svc.NewServiceContext(c)
	logx.Must(register.RegKq(group, ctx))

	// Asynq不加入ServiceGroup
	logx.Must(register.RegDq(ctx))

	fmt.Printf("Starting service: %s\n", c.Name)
	group.Start()
}
