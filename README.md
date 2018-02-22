## Athena Back-End Server

此为 Athena 项目的后端服务。主要承载自 API Server 的请求，沟通数据库，返回请求。通过 gRPC 和 API Server 交互。

## 其他

数据库采用了 postgreSQL，个人感觉比较稳。

`proto` 文件和 API Server 共用。

## env

release

## 注意

尚未部署至生产环境

## cells 中 premission 字段(单词拼错了，应该是 permission)
-- 2017-10-06 添加权限控制字段
-- 2  -> public 共有的，谁都可以看
-- 3  -> 受保护的，只有发布者自己能看
-- 4+ -> 暂未定义

## users 中 role 字段