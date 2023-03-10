// Code generated by goctl. DO NOT EDIT.
// Source: comment.proto

package server

import (
	"context"

	"OutTiktok/apps/comment/comment"
	"OutTiktok/apps/comment/internal/logic"
	"OutTiktok/apps/comment/internal/svc"
)

type CommentServer struct {
	svcCtx *svc.ServiceContext
	comment.UnimplementedCommentServer
}

func NewCommentServer(svcCtx *svc.ServiceContext) *CommentServer {
	return &CommentServer{
		svcCtx: svcCtx,
	}
}

func (s *CommentServer) Action(ctx context.Context, in *comment.ActionReq) (*comment.ActionRes, error) {
	l := logic.NewActionLogic(ctx, s.svcCtx)
	return l.Action(in)
}

func (s *CommentServer) List(ctx context.Context, in *comment.ListReq) (*comment.ListRes, error) {
	l := logic.NewListLogic(ctx, s.svcCtx)
	return l.List(in)
}

func (s *CommentServer) GetCommentCount(ctx context.Context, in *comment.GetCommentCountReq) (*comment.GetCommentCountRes, error) {
	l := logic.NewGetCommentCountLogic(ctx, s.svcCtx)
	return l.GetCommentCount(in)
}
