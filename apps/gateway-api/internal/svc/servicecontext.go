// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"github.com/Yooz-1999/agentforge/apps/core-rpc/core"
	"github.com/Yooz-1999/agentforge/apps/gateway-api/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Core   core.Core
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Core:   core.NewCore(zrpc.MustNewClient(c.CoreRPC)),
	}
}
