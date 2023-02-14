package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf
	Consul   consul.Conf
	MysqlDsn string
	Minio    struct {
		Host        string
		AccessKey   string
		SecretKey   string
		VideoBucket string
		CoverBucket string
	}
	User zrpc.RpcClientConf
}
