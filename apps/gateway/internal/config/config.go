package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	User     zrpc.RpcClientConf
	Publish  zrpc.RpcClientConf
	Favorite zrpc.RpcClientConf
	Comment  zrpc.RpcClientConf
	Relation zrpc.RpcClientConf
	Feed     zrpc.RpcClientConf
}
