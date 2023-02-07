package logic

import (
	"context"

	"OutTiktok/apps/feed/feed"
	"OutTiktok/apps/feed/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeedLogic) Feed(in *feed.FeedReq) (*feed.FeedRes, error) {
	// todo: add your logic here and delete this line

	return &feed.FeedRes{}, nil
}
