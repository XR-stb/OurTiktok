package svc

import (
	"OutTiktok/apps/comment/internal/config"
	"OutTiktok/apps/user/user"
	"OutTiktok/apps/user/userclient"
	"OutTiktok/dao"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config     config.Config
	DB         *gorm.DB
	UserClient user.UserClient
	UserCache  map[int64]*userclient.UserInfo
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		DB:         dao.NewGorm(c.MysqlDsn),
		UserClient: userclient.NewUser(zrpc.MustNewClient(c.User, zrpc.WithDialOption(grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`)))),
		UserCache:  map[int64]*userclient.UserInfo{},
	}
}
