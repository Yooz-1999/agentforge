// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"github.com/Yooz-1999/agentforge/apps/chat-api/internal/config"
	"github.com/Yooz-1999/agentforge/apps/chat-api/internal/middleware"
	"github.com/Yooz-1999/agentforge/apps/core-rpc/core"
	"github.com/Yooz-1999/agentforge/pkg/auth"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	Core       core.Core
	AccessAuth rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	tokenManager := auth.NewTokenManager(
		c.JWT.AccessSecret,
		c.JWT.RefreshSecret,
		c.JWT.AccessExpireSeconds,
		c.JWT.RefreshExpireSeconds,
	)

	return &ServiceContext{
		Config:     c,
		Core:       core.NewCore(zrpc.MustNewClient(c.CoreRPC)),
		AccessAuth: middleware.NewAccessAuthMiddleware(tokenManager).Handle,
	}
}
