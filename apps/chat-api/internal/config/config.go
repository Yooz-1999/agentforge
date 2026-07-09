// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package config

import (
	sharedconfig "github.com/Yooz-1999/agentforge/pkg/config"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	CoreRPC zrpc.RpcClientConf
	JWT     sharedconfig.JWTConf
	LLM     sharedconfig.LLMConf
}
