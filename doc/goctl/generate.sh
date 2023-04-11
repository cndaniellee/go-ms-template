

#生成Api部分代码（在根目录执行，*改为模块名称）
goctl api go --api ./doc/goctl/api/*.api --dir ./service/*/api --home ./doc/goctl/template --style go_zero


#生成Rpc部分代码（在根目录执行，*改为模块名称）
goctl rpc protoc ./doc/goctl/rpc/*.proto --zrpc_out=./service/*/rpc --go_out=./service/*/rpc/pb --go-grpc_out=./service/*/rpc/pb --home ./doc/goctl/template --style go_zero


#生成Dockerfile（在模块目录中执行，*改为模块名称）
goctl docker --go *.go --exe run


#构建Docker镜像
docker build -t user-api:latest -f service/user/api/Dockerfile .


#生成K8S配置文件
goctl kube deploy -name user-api -namespace goms -image 192.168.2.220:8443/goms/user-api:1.0.0 -o kube.yaml -port 7801 -port 9801