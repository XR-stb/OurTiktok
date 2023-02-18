package svc

import (
	"OutTiktok/apps/feed/internal/config"
	"OutTiktok/apps/publish/publish"
	"OutTiktok/apps/publish/publishclient"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type ServiceContext struct {
	Config        config.Config
	Redis         *redis.Redis
	PublishClient publish.PublishClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		Redis:         redis.New(c.Redis.Host),
		PublishClient: publishclient.NewPublish(zrpc.MustNewClient(c.Publish, zrpc.WithDialOption(grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`)))),
	}
}
