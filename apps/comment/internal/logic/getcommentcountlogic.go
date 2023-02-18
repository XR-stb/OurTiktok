package logic

import (
	"context"
	"fmt"

	"OutTiktok/apps/comment/comment"
	"OutTiktok/apps/comment/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentCountLogic {
	return &GetCommentCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentCountLogic) GetCommentCount(in *comment.GetCommentCountReq) (*comment.GetCommentCountRes, error) {
	counts := make([]int64, len(in.VideoIds))
	for i, id := range in.VideoIds {
		key := fmt.Sprintf("cids_%d", id)
		count, _ := l.svcCtx.Redis.Zcard(key)
		counts[i] = int64(count)
	}

	return &comment.GetCommentCountRes{
		Counts: counts,
	}, nil
}
