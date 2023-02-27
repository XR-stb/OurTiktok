package svc

import (
	"OutTiktok/apps/comment/comment"
	"OutTiktok/apps/comment/commentclient"
	"OutTiktok/apps/favorite/favorite"
	"OutTiktok/apps/favorite/favoriteclient"
	"OutTiktok/apps/feed/feed"
	"OutTiktok/apps/feed/feedclient"
	"OutTiktok/apps/gateway/internal/config"
	"OutTiktok/apps/message/message"
	"OutTiktok/apps/message/messageclient"
	"OutTiktok/apps/publish/publish"
	"OutTiktok/apps/publish/publishclient"
	"OutTiktok/apps/relation/relation"
	"OutTiktok/apps/relation/relationclient"
	"OutTiktok/apps/user/user"
	"OutTiktok/apps/user/userclient"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type ServiceContext struct {
	Config         config.Config
	UserClient     user.UserClient
	PublishClient  publish.PublishClient
	FavoriteClient favorite.FavoriteClient
	CommentClient  comment.CommentClient
	RelationClient relation.RelationClient
	FeedClient     feed.FeedClient
	MessageClient  message.MessageClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UserClient:     userclient.NewUser(zrpc.MustNewClient(c.User, zrpc.WithDialOption(grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`)))),
		PublishClient:  publishclient.NewPublish(zrpc.MustNewClient(c.Publish, zrpc.WithDialOption(grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`)))),
		FavoriteClient: favoriteclient.NewFavorite(zrpc.MustNewClient(c.Favorite, zrpc.WithDialOption(grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`)))),
		CommentClient:  commentclient.NewComment(zrpc.MustNewClient(c.Comment, zrpc.WithDialOption(grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`)))),
		RelationClient: relationclient.NewRelation(zrpc.MustNewClient(c.Relation, zrpc.WithDialOption(grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`)))),
		FeedClient:     feedclient.NewFeed(zrpc.MustNewClient(c.Feed, zrpc.WithDialOption(grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`)))),
		MessageClient:  messageclient.NewMessage(zrpc.MustNewClient(c.Message, zrpc.WithDialOption(grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`)))),
	}
}
