# agentforge
Forge Your AI Agents.

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

公共配置结构统一放在 [pkg/config/common.go](/Users/zichen/Documents/go/src/agentforge/pkg/config/common.go:1)。

规则是：

1. JWT、MySQL、Redis、模型平台配置这类公共字段，只在公共层定义一份
2. 每个服务只保留自己独有的配置，比如 `RestConf` 或 `RpcServerConf`
3. 后面新增服务时，直接复用公共配置结构，不要再复制一套新的
