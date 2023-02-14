package svc

import (
	"OutTiktok/apps/favorite/favorite"
	"OutTiktok/apps/favorite/favoriteclient"
	"OutTiktok/apps/publish/internal/config"
	"OutTiktok/apps/publish/pkg/snowflake"
	"OutTiktok/apps/user/user"
	"OutTiktok/apps/user/userclient"
	"OutTiktok/dao"
	"github.com/minio/minio-go/v6"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config         config.Config
	Minio          *minio.Client
	DB             *gorm.DB
	Sf             *snowflake.Snowflake
	UserClient     user.UserClient
	FavoriteClient favorite.FavoriteClient
	UserCache      map[int64]*userclient.UserInfo
}

func NewServiceContext(c config.Config) *ServiceContext {
	minioClient, err := minio.New(c.Minio.Host, c.Minio.AccessKey, c.Minio.SecretKey, false)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:         c,
		Minio:          minioClient,
		DB:             dao.NewGorm(c.MysqlDsn),
		Sf:             &snowflake.Snowflake{},
		UserClient:     userclient.NewUser(zrpc.MustNewClient(c.User, zrpc.WithDialOption(grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`)))),
		FavoriteClient: favoriteclient.NewFavorite(zrpc.MustNewClient(c.Favorite, zrpc.WithDialOption(grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`)))),
		UserCache:      map[int64]*userclient.UserInfo{},
	}
}
