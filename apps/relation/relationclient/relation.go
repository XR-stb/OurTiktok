// Code generated by goctl. DO NOT EDIT.
// Source: relation.proto

package relationclient

import (
	"context"

	"OutTiktok/apps/relation/relation"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ActionReq       = relation.ActionReq
	ActionRes       = relation.ActionRes
	FollowListReq   = relation.FollowListReq
	FollowListRes   = relation.FollowListRes
	FollowerListReq = relation.FollowerListReq
	FollowerListRes = relation.FollowerListRes
	FriendListReq   = relation.FriendListReq
	FriendListRes   = relation.FriendListRes
	GetRelationsReq = relation.GetRelationsReq
	GetRelationsRes = relation.GetRelationsRes
	UserInfo        = relation.UserInfo

	Relation interface {
		Action(ctx context.Context, in *ActionReq, opts ...grpc.CallOption) (*ActionRes, error)
		FollowList(ctx context.Context, in *FollowListReq, opts ...grpc.CallOption) (*FollowListRes, error)
		FollowerList(ctx context.Context, in *FollowerListReq, opts ...grpc.CallOption) (*FollowerListRes, error)
		FriendList(ctx context.Context, in *FriendListReq, opts ...grpc.CallOption) (*FriendListRes, error)
		GetRelations(ctx context.Context, in *GetRelationsReq, opts ...grpc.CallOption) (*GetRelationsRes, error)
	}

	defaultRelation struct {
		cli zrpc.Client
	}
)

func NewRelation(cli zrpc.Client) Relation {
	return &defaultRelation{
		cli: cli,
	}
}

func (m *defaultRelation) Action(ctx context.Context, in *ActionReq, opts ...grpc.CallOption) (*ActionRes, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.Action(ctx, in, opts...)
}

func (m *defaultRelation) FollowList(ctx context.Context, in *FollowListReq, opts ...grpc.CallOption) (*FollowListRes, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.FollowList(ctx, in, opts...)
}

func (m *defaultRelation) FollowerList(ctx context.Context, in *FollowerListReq, opts ...grpc.CallOption) (*FollowerListRes, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.FollowerList(ctx, in, opts...)
}

func (m *defaultRelation) FriendList(ctx context.Context, in *FriendListReq, opts ...grpc.CallOption) (*FriendListRes, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.FriendList(ctx, in, opts...)
}

func (m *defaultRelation) GetRelations(ctx context.Context, in *GetRelationsReq, opts ...grpc.CallOption) (*GetRelationsRes, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.GetRelations(ctx, in, opts...)
}