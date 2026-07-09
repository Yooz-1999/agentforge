package config

import (
	sharedconfig "github.com/Yooz-1999/agentforge/pkg/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	MySQL sharedconfig.MySQLConf
	Redis sharedconfig.RedisConf
}
