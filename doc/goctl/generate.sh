

#生成Api部分代码（在当前目录执行，*改为模块名称）
goctl api go --api ./doc/goctl/api/order.api --dir ./service/order/api --home ./doc/goctl/template --style go_zero


goctl rpc protoc ./doc/goctl/rpc/order.proto --zrpc_out=./service/order/rpc --go_out=./service/order/rpc/pb --go-grpc_out=./service/order/rpc/pb --home ./doc/goctl/template --style go_zero


#生成Dockerfile（在模块目录中执行，*改为模块名称）
goctl docker --go *.go --exe run