package feed

import (
	"OutTiktok/apps/feed/feed"
	"OutTiktok/apps/gateway/internal/svc"
	"OutTiktok/apps/gateway/internal/types"
	"OutTiktok/apps/gateway/pkg/jwt"
	"context"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedLogic) Feed(req *types.FeedReq) (resp *types.FeedRes, err error) {
	resp = &types.FeedRes{}

	// 验证Token
	var userId int64
	claims, err := jwt.VerifyToken(req.Token)
	if err == nil {
		userId = claims.UserId
	}

	r, err := l.svcCtx.FeedClient.Feed(context.Background(), &feed.FeedReq{
		UserId:     userId,
		LatestTime: req.LatestTime,
	})
	if err != nil || r.Status != 0 {
		resp.StatusCode = -1
		resp.StatusMsg = "获取失败"
		return
	}

	resp.NextTime = r.NextTime
	resp.VideoList = make([]types.Video, len(r.VideoList))
	for i, video := range r.VideoList {
		_ = copier.Copy(&resp.VideoList[i], &video)
	}

	return
}
