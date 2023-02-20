package svc

import (
	"OutTiktok/apps/favorite/favorite"
	"OutTiktok/apps/favorite/favoriteclient"
	"OutTiktok/apps/publish/publish"
	"OutTiktok/apps/publish/publishclient"
	"OutTiktok/apps/relation/relation"
	"OutTiktok/apps/relation/relationclient"
	"OutTiktok/apps/user/internal/config"
	"OutTiktok/dao"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config         config.Config
	DB             *gorm.DB
	Redis          *redis.Redis
	PublishClient  publish.PublishClient
	FavoriteClient favorite.FavoriteClient
	RelationClient relation.RelationClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		DB:             dao.NewGorm(c.MysqlDsn),
		Redis:          redis.New(c.Redis.Host),
		PublishClient:  publishclient.NewPublish(zrpc.MustNewClient(c.Publish, zrpc.WithDialOption(grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`)))),
		FavoriteClient: favoriteclient.NewFavorite(zrpc.MustNewClient(c.Favorite, zrpc.WithDialOption(grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`)))),
		RelationClient: relationclient.NewRelation(zrpc.MustNewClient(c.Relation, zrpc.WithDialOption(grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`)))),
	}
}
