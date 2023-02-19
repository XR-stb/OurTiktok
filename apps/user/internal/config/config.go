package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf
	Consul   consul.Conf
	MysqlDsn string
	Redis    redis.RedisConf
	Minio    struct {
		Host string
	}
}
