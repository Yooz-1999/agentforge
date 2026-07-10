# AgentForge 项目记忆

## 项目定位

已确认：

1. AgentForge 是一个 AI Agent 平台后端项目。
2. 产品目标是支持创建 AI Agent、Workflow、工具调用和多模型接入。
3. 用户登录后可以创建 Agent，通过聊天调用 Agent；Agent 后续会调用工具、执行工作流、返回流式结果。
4. 当前仓库只放后端代码，前端建议另起仓库，例如 `agentforge-web`。

## 技术架构

已确认：

1. 后端使用 Go。
2. 微服务框架使用 go-zero。
3. 当前服务包括 `core-rpc`、`gateway-api`、`chat-api`。
4. 公共配置结构放在 `pkg/config`，各服务只保留自己独有的启动配置。
5. `go.mod` 和 `go.sum` 应该提交到远程仓库，不应该加入 `.gitignore`。

## 本机开发环境

已确认：

1. 用户当前更倾向使用本机 MySQL 和 Redis，不想依赖 Docker，因为 Docker 占用资源较多。
2. 本机 MySQL 已安装并运行，版本曾确认为 MySQL 9.6.0。
3. 本机 MySQL 可以使用 `root` 无密码连接 `127.0.0.1:3306`。
4. 本机 Redis 已安装并运行，`redis-cli ping` 曾返回正常。
5. 数据库名是 `agentforge`。
6. 当前已建表：`users`、`agents`、`conversations`、`messages`、`user_api_keys`。
7. 数据库排序规则可以使用 `utf8mb4_0900_ai_ci`，如果没有则 `utf8mb4_unicode_520_ci` 也可以。

注意：

1. 不要把 Docker 示例里的 `root:password` 当成本机 MySQL 默认配置。
2. 如果用户后续改了 MySQL 密码，应以本机真实可登录配置为准。
3. 不要把真实密码写进项目级 skill。

## 本机服务启动

当前本机开发模式已确认不强制依赖 `etcd`。

启动顺序：

```bash
cd /Users/zichen/Documents/go/src/agentforge/apps/core-rpc
go run core.go -f etc/core.yaml
```

```bash
cd /Users/zichen/Documents/go/src/agentforge/apps/gateway-api
go run gateway.go -f etc/gateway-api.yaml
```

```bash
cd /Users/zichen/Documents/go/src/agentforge/apps/chat-api
go run chat.go -f etc/chat-api.yaml
```

端口：

1. `core-rpc` 监听 `0.0.0.0:8080`
2. `gateway-api` 监听 `0.0.0.0:8888`
3. `chat-api` 监听 `0.0.0.0:8889`

服务连接方式：

1. `gateway-api` 本机直接连接 `127.0.0.1:8080`
2. `chat-api` 本机直接连接 `127.0.0.1:8080`
3. 本机启动不需要先启动 `etcd`

## 已实现接口状态

已确认可用：

1. `POST /api/v1/auth/register`
2. `POST /api/v1/auth/login`
3. `POST /api/v1/auth/refresh`

已确认行为：

1. 注册会检查邮箱、密码、昵称。
2. 注册会检查邮箱格式。
3. 注册会要求密码至少 6 位。
4. 注册会检查邮箱是否已存在。
5. 注册会加密密码后写入 `users` 表。
6. 登录成功会生成 access token 和 refresh token。
7. refresh token 会写入 Redis，key 格式为 `auth:refresh:{user_id}`。
8. access token 和 refresh token 使用不同的签名密钥，并检查各自的 token 类型。
9. Agent、Conversation、Message、Chat Stream 路由已经接入 access token 鉴权；这些业务本身仍未实现。
10. 刷新令牌前会重新查询数据库，只允许未删除且状态正常的用户刷新；用户失效时会删除 Redis 中的旧刷新令牌。

未完成：

1. Agent 相关接口路由已有，但业务未实现。
2. Conversation 相关接口路由已有，但业务未实现。
3. Message 相关接口路由已有，但业务未实现。
4. Chat Stream 路由已有，但业务未实现。
5. 未实现接口应该明确返回 `接口还没有实现`，不要返回空成功结果。

## 代码约定

用户偏好：

1. README 和设计文档使用中文。
2. repository 方法需要中文注释，方便开发人员理解。
3. 最终沟通要结果优先，解释清楚发生了什么、为什么、怎么处理。
4. 排查问题时按概率从高到低说明。

实现约定：

1. 优先复用已有公共结构、标准库和依赖能力。
2. 只有现有能力不适合时才新增自定义实现。
3. 本机配置里业务 Redis 字段使用 `AppRedis`，避免和 go-zero RPC 内部 `Redis` 字段重名。
4. `core-rpc` 当前不需要业务 Redis 配置，不要为了看起来统一而加回去。
5. `gateway-api` 和 `chat-api` 不使用 go-zero 自带的 HTTP 请求内容日志，因为服务端错误时该日志会记录完整请求头和请求体，可能暴露密码或令牌。
6. `core-rpc` 必须禁止记录 `RegisterUser` 和 `VerifyLogin` 的请求内容。

## 已踩过的坑

1. `go-zero` 的 RPC 配置里有内部 `Redis` 字段；如果业务配置也叫 `Redis`，可能触发 `field "redis.Host" is not set`。
2. `core-rpc` 原来写 `root:password` 会导致本机 MySQL 报 `Access denied for user 'root'@'localhost'`。
3. 旧的 `bitnami/etcd:3.5` 镜像无法拉取，Docker compose 已改用 `registry.k8s.io/etcd:3.5.14-0`。
4. 未实现接口如果返回空成功，会误导排查；应该明确返回未实现错误。
5. go-zero 默认 RPC 日志会记录完整注册和登录请求，必须把这两个方法加入内容日志排除列表。
6. go-zero 默认 HTTP 日志在 500/501 响应时会记录完整请求，鉴权失败日志还会记录 Authorization 请求头；不能直接用于包含密码和令牌的接口。
7. 本机 MySQL 使用无密码配置；Docker MySQL 使用 `core-docker.yaml`，不要混用两套配置。

## 验证记录

已确认通过：

1. `go test ./...`
2. `go vet ./...`
3. 三个服务可启动并监听端口。
4. 注册接口可写入用户。
5. 登录接口可返回 token。
6. 测试用户和 Redis 测试缓存已清理。
7. access token 可通过业务路由鉴权，refresh token 和无令牌请求会被拒绝。
8. 用户禁用后刷新令牌会被拒绝，Redis 旧令牌会被删除。
9. 注册登录日志已确认不再包含密码，HTTP 日志也不再记录 access token 和 refresh token。
10. 已为令牌类型、签名算法、鉴权中间件、中文输入长度和 bcrypt 密码上限增加自动化测试。
