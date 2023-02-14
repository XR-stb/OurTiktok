package logic

import (
	"context"

	"OutTiktok/apps/comment/comment"
	"OutTiktok/apps/comment/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *comment.ListReq) (*comment.ListRes, error) {

	return &comment.ListRes{}, nil
}
