

#生成Api部分代码（在当前目录执行，*改为模块名称）
goctl api go --api ./doc/goctl/api/product.api --dir ./app/product/api --home ./doc/goctl/template --style go_zero


goctl rpc protoc ./doc/goctl/rpc/product.proto --zrpc_out=./app/product/rpc --go_out=./app/product/rpc/pb --go-grpc_out=./app/product/rpc/pb --home ./doc/goctl/template --style go_zero


#生成Dockerfile（在模块目录中执行，*改为模块名称）
goctl docker --go *.go --exe run