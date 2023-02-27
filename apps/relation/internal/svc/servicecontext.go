package svc

import (
	"OutTiktok/apps/message/messageclient"
	"OutTiktok/apps/relation/internal/config"
	"OutTiktok/apps/user/user"
	"OutTiktok/apps/user/userclient"
	"OutTiktok/dao"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config        config.Config
	DB            *gorm.DB
	Redis         *redis.Redis
	UserClient    userclient.User
	MessageClient messageclient.Message
	UserCache     map[int64]*user.UserInfo
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		DB:            dao.NewGorm(c.MysqlDsn),
		Redis:         redis.New(c.Redis.Host),
		UserClient:    userclient.NewUser(zrpc.MustNewClient(c.User, zrpc.WithDialOption(grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`)))),
		MessageClient: messageclient.NewMessage(zrpc.MustNewClient(c.Message, zrpc.WithDialOption(grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`)))),
		UserCache:     map[int64]*user.UserInfo{},
	}
}
