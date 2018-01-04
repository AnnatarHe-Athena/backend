## Athena Back-End Server

此为 Athena 项目的后端服务。主要承载自 API Server 的请求，沟通数据库，返回请求。通过 gRPC 和 API Server 交互。

## 其他

数据库采用了 postgreSQL，个人感觉比较稳。

`proto` 文件和 API Server 共用。

## env

release

## 注意

尚未部署至生产环境