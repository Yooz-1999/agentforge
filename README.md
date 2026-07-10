# agentforge
一个支持创建 AI Agent、工作流、工具调用和多模型接入的后端项目。

## 仓库范围

这个仓库是 AgentForge 的后端工程仓库。

这里会放：

1. `gateway-api`
2. `chat-api`
3. `core-rpc`
4. 后端公共配置、脚本、部署文件、SQL 迁移脚本

这里不会放前端页面代码。

前端建议放在单独仓库里，例如 `agentforge-web`。

## 配置策略

公共配置结构统一放在 [pkg/config/common.go](pkg/config/common.go)。

规则是：

1. JWT、MySQL、Redis、模型平台配置这类公共字段，只在公共层定义一份
2. 每个服务只保留自己独有的配置，比如 `RestConf` 或 `RpcServerConf`
3. 后面新增服务时，直接复用公共配置结构，不要再复制一套新的

## 项目记忆

项目级 skill 放在 [skills/agentforge-project/SKILL.md](skills/agentforge-project/SKILL.md)。

这里用来记录后续对话里值得长期保留的项目事实、启动方式、避坑经验和开发约定。

## 本机开发

当前项目支持直接使用本机安装的 MySQL 和 Redis 开发，不强制依赖 Docker。

本机默认启动方式也不再强制依赖 `etcd`。

这意味着：

1. `core-rpc` 会直接监听 `127.0.0.1:8080`
2. `gateway-api` 会直接连接 `127.0.0.1:8080`
3. 你本地只要先准备好 MySQL 和 Redis，就可以先把注册登录链路跑起来

建议做法是：

1. 本机启动 MySQL
2. 本机启动 Redis
3. 创建数据库 `agentforge`
4. 执行 [sql/migrations/001_init.sql](sql/migrations/001_init.sql) 建表

数据库创建时建议：

1. 数据库名：`agentforge`
2. 字符集：`utf8mb4`
3. 排序规则：优先 `utf8mb4_0900_ai_ci`
4. 如果你的本机版本里没有 `utf8mb4_0900_ai_ci`，用 `utf8mb4_unicode_520_ci` 也可以

说明：

1. 你截图里选的 `utf8mb4_unicode_520_ci` 可以正常用
2. 如果本机是 MySQL 8/9，一般更推荐 `utf8mb4_0900_ai_ci`
3. 项目配置里的 MySQL 账号密码，应该改成你本机实际能登录的那组，不要默认照搬 Docker 示例值

## 启动后端服务

启动前先检查一个地方：

1. 打开 [apps/core-rpc/etc/core.yaml](apps/core-rpc/etc/core.yaml)
2. 默认配置使用 `root` 无密码连接本机 MySQL
3. 如果你的 MySQL 设置了密码，把 `MySQL.DataSource` 改成你本机真正能登录的账号密码

然后按这个顺序启动：

1. 启动 `core-rpc`

```bash
go run ./apps/core-rpc -f apps/core-rpc/etc/core.yaml
```

2. 启动 `gateway-api`

```bash
go run ./apps/gateway-api -f apps/gateway-api/etc/gateway-api.yaml
```

如果两个服务都正常启动，你会看到：

1. `core-rpc` 监听 `0.0.0.0:8080`
2. `gateway-api` 监听 `0.0.0.0:8888`
3. `chat-api` 本机开发时也会直接连接 `127.0.0.1:8080`，不再强制依赖 `etcd`

### 使用 Docker 启动 MySQL 和 Redis

Docker 里的 MySQL root 密码是 `password`，和本机无密码配置不同。

从仓库根目录执行：

```bash
docker compose -f deploy/docker-compose/local-infra.yml up -d mysql redis
go run ./apps/core-rpc -f apps/core-rpc/etc/core-docker.yaml
```

使用本机 MySQL 时继续使用 `etc/core.yaml`，不要改成 Docker 配置。

## 当前已可用接口

目前已经有下面这几个接口：

1. `POST /api/v1/auth/register`
2. `POST /api/v1/auth/login`
3. `POST /api/v1/auth/refresh`

Agent、会话、消息和聊天接口需要在 `Authorization` 请求头中携带 access token。refresh token 只能用于刷新接口，不能调用这些业务接口。

仓库内的 JWT 密钥只用于本机开发。部署到测试或生产环境前，必须分别替换 access token 和 refresh token 的密钥，不能继续使用 `change-me-access`、`change-me-refresh`。

`gateway-api` 和 `chat-api` 当前关闭了 go-zero 自带的 HTTP 请求内容日志，因为接口返回 500 或 501 时，框架会把请求头和请求体完整写入日志，可能暴露密码或令牌。`core-rpc` 也明确禁止记录注册和登录请求内容。后续如果重新开启 HTTP 请求日志，必须先确认日志只记录请求方法、路径、状态码和耗时。

### 注册接口怎么测

```bash
curl -X POST http://127.0.0.1:8888/api/v1/auth/register \
  -H 'Content-Type: application/json' \
  -d '{
    "email":"test@example.com",
    "password":"123456",
    "nickname":"test-user"
  }'
```

注册时系统会做这些事：

1. 检查邮箱、密码、昵称是不是空
2. 检查邮箱格式对不对
3. 检查密码长度是不是至少 6 位
4. 先去 `users` 表查这个邮箱是否已经存在
5. 如果不存在，就把密码加密后写入 `users` 表
6. 最后返回新用户的 `id`、`email` 和 `nickname`

### 登录接口怎么测

```bash
curl -X POST http://127.0.0.1:8888/api/v1/auth/login \
  -H 'Content-Type: application/json' \
  -d '{
    "email":"test@example.com",
    "password":"123456"
  }'
```
