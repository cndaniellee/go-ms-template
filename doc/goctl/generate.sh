

#生成Api部分代码（在当前目录执行，*改为模块名称）
goctl api go --api ./api/*.api --dir ../../app/*/api --home ./template --style go_zero

#生成Dockerfile（在模块目录中执行，*改为模块名称）
goctl docker --go *.go --exe run