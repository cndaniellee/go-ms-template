## go-ms-template

English | [简体中文](README-zhCN.md)

#### Microservice Development Template (Base go-zero)

I came into contact with Golang by some chance after I joined a company under Ant in 2018. At that time, I was only used to write internal tool applications. I selected iris, which is still growing, from beego, gin, echo and iris as the basic framework. \
\
However, Golang gradually replaced Java as my primary development language (really sweet), so naturally I needed to develop large scale applications. Java microservices side of the spring family bucket is sitting on a pedestal, while Golang side is in a state of a hundred flowers bloom. \
\
From go-kit in the early toolkit form to go-micro, go-zero, kratos, kitex, etc., Go-Zero was the first framework I got, for the same reasons as when I chose iris: community, contribution activity, ecology, and tools (goctl etc.) all felt good. \
\
In order to facilitate the follow-up of new projects, I integrate the tools that may be used into a template, which includes some simple shop apis, and open source it for everyone to learn and use. Welcome Star/Issue, and jointly improve the quality of our code.

#### Tool list

- Docker - Container Management (Dev ENV)
- Air - Hot load (Dev ENV)
- K8S - Service Cluster (Pro ENV)
- Docker Compose - Deployment tool
- Go-Zero - API/RPC framework
- Validator.V9 - Parameter verification
- Nginx - Reverse proxy
- Etcd - Registry
- DTM - Distributed transaction
- Redis - Distributed cache/lock
- RedisManager - Redis Monitors
- MySQL - Database
- Zookeeper-Registry (Kafka only)
- Kafka - Message queue
- Asynq - Delay queue
- AsynqMon-Asynq Monitors
- ElasticSearch - Storage/search induced
- Kibana-ES View/analyzer
- FileBeat - Collects logs
- GoStash - Log consumption
- Jeager - Link tracing
- Prometheus - Container monitoring
- Grafana - Visualization
- Jenkins - CI/CD
- Gitlab - Private warehouse
- Harbor - Mirror warehouse

#### Start-up project

If you already have Docker installed, this project can be launched directly: \
- Start dependency tool \
`[root@dev go-ms-template]# docker compose -f docker-compose-env.yml up -d`
- Launch project application \
- `[root@dev go-ms-template]# docker compose -f docker-compose.yml up -d`

#### directory structure

The directory structure of this project is very concise, and high readability is my constant pursuit.

- common: common components (login, log, message, return, database, parameter verification, etc.)

- doc: indicates the document content

  - conf: depends on the configuration file of the tool
  - dtm: initializes the SQL of the DTM database and copies it
  - goctl: defines the api, proto file, modified goctl template, and the goctl command line

- service: indicates the project application

  - XXX: indicates the application module

    - api: indicates the HTTP service of the module
    - mq: indicates the message queue or delay queue consumption end of the module
    - rpc: RPC service of the module

#### Component architecture
![architecture.jpg](doc%2Farchitecture.jpg)

#### Thank @Mikaelemmmm for the inspiration given by part of the project.