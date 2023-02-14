package svc

import (
	"OutTiktok/apps/gateway/internal/config"
	"OutTiktok/apps/publish/publish"
	"OutTiktok/apps/publish/publishclient"
	"OutTiktok/apps/user/user"
	"OutTiktok/apps/user/userclient"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type ServiceContext struct {
	Config        config.Config
	UserClient    user.UserClient
	PublishClient publish.PublishClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserClient:    userclient.NewUser(zrpc.MustNewClient(c.User, zrpc.WithDialOption(grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`)))),
		PublishClient: publishclient.NewPublish(zrpc.MustNewClient(c.Publish, zrpc.WithDialOption(grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`)))),
	}
}
