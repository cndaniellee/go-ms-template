## go-ms-template

[English](README.md) | 简体中文

#### 微服务开发模板（Base go-zero）
我在2018年入职一家蚂蚁旗下公司后的某种机缘巧合中接触到了Golang，那时候只用来写内部工具型应用，我从beego、gin、echo、iris中选出了还在成长中的iris作为基础框架。 \
\
然而后来Golang慢慢取代Java成为了我的主要开发语言（真的香），自然需要向开发大型应用的方向研究。Java微服务那边spring全家桶是稳坐高台的，而Golang这边却处于百花齐放难寻一枝的状态。 \
\
从早期工具包形态的go-kit，后有go-micro、go-zero、kratos、kitex等等，我最先拿到手的就是Go-Zero，理由与我选择iris时相似，不管是社区、贡献活跃度还是生态、工具（goctl香）都感觉很不错。 \
\
为了方便后续起新项目，我把可能用到的工具整合成了一个模板，并包含了一套简单的店铺功能，开源出来供大家学习使用，欢迎Star/Issue，共同提升我们的代码质量。

#### 工具列表

- Docker - 容器管理 (Dev ENV)
- Air - 热加载 (Dev ENV)
- K8S - 服务集群 (Pro ENV)
- Docker Compose - 部署工具
- Go-Zero - API/RPC框架
- Validator.V9 - 参数校验
- Nginx - 反向代理
- Etcd - 注册中心
- DTM - 分布式事务
- Redis - 分布式缓存/锁
- RedisManager - Redis监控
- MySQL - 数据库
- Zookeeper - 注册中心（仅Kafka）
- Kafka - 消息队列
- Asynq - 延迟队列
- AsynqMon - Asynq监控
- ElasticSearch - 存储/搜索引起
- Kibana - ES查看/分析器
- FileBeat - 日志收集
- GoStash - 日志消费
- Jeager - 链路追踪
- Prometheus - 容器监控
- Grafana - 可视化
- Jenkins - CI/CD
- Gitlab - 私有仓库
- Harbor - 镜像仓库

#### 启动项目

如果你已经安装好了Docker，本项目可以直接启动： \
- 启动依赖工具 \
`[root@dev go-ms-template]# docker compose -f docker-compose-env.yml up -d`
- 启动项目应用 \
- `[root@dev go-ms-template]# docker compose -f docker-compose.yml up -d`


#### 目录结构
本项目的目录结构整理的非常简洁，高可读性是我的持续追求。

- common：通用组件（登录、日志、消息、返回、数据库、参数校验等）

- doc：文档内容

    - conf：依赖工具的配置文件
    - dtm：初始化DTM数据库的SQL，复制了一份过来
    - goctl：api、proto文件的定义，以及修改过的goctl模板、执行goctl的命令行

- service：项目应用

    - XXX: 应用模块
  
        - api：模块的HTTP服务
        - mq：模块的消息队列、延迟队列消费端
        - rpc：模块的RPC服务

#### 组件架构
![architecture.jpg](doc%2Farchitecture.jpg)

#### 感谢 @Mikaelemmmm 项目中部分内容给予的启发。
