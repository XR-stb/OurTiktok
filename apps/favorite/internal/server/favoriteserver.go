// Code generated by goctl. DO NOT EDIT.
// Source: favorite.proto

package server

import (
	"context"

	"OutTiktok/apps/favorite/favorite"
	"OutTiktok/apps/favorite/internal/logic"
	"OutTiktok/apps/favorite/internal/svc"
)

type FavoriteServer struct {
	svcCtx *svc.ServiceContext
	favorite.UnimplementedFavoriteServer
}

func NewFavoriteServer(svcCtx *svc.ServiceContext) *FavoriteServer {
	return &FavoriteServer{
		svcCtx: svcCtx,
	}
}

func (s *FavoriteServer) Action(ctx context.Context, in *favorite.ActionReq) (*favorite.ActionRes, error) {
	l := logic.NewActionLogic(ctx, s.svcCtx)
	return l.Action(in)
}

func (s *FavoriteServer) List(ctx context.Context, in *favorite.ListReq) (*favorite.ListRes, error) {
	l := logic.NewListLogic(ctx, s.svcCtx)
	return l.List(in)
}
