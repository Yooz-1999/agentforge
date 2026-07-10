// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"github.com/Yooz-1999/agentforge/apps/core-rpc/core"
	"github.com/Yooz-1999/agentforge/apps/gateway-api/internal/config"
	"github.com/Yooz-1999/agentforge/apps/gateway-api/internal/middleware"
	"github.com/Yooz-1999/agentforge/pkg/auth"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	Core         core.Core
	Redis        *redis.Client
	TokenManager *auth.TokenManager
	AccessAuth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	zrpc.DontLogClientContentForMethod("/core.Core/RegisterUser")
	zrpc.DontLogClientContentForMethod("/core.Core/VerifyLogin")
	tokenManager := auth.NewTokenManager(
		c.JWT.AccessSecret,
		c.JWT.RefreshSecret,
		c.JWT.AccessExpireSeconds,
		c.JWT.RefreshExpireSeconds,
	)

	return &ServiceContext{
		Config: c,
		Core:   core.NewCore(zrpc.MustNewClient(c.CoreRPC)),
		Redis: redis.NewClient(&redis.Options{
			Addr:     c.AppRedis.Addr,
			Password: c.AppRedis.Password,
			DB:       c.AppRedis.DB,
		}),
		TokenManager: tokenManager,
		AccessAuth:   middleware.NewAccessAuthMiddleware(tokenManager).Handle,
	}
}
