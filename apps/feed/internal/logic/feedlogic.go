package logic

import (
	"OutTiktok/apps/feed/feed"
	"OutTiktok/apps/feed/internal/svc"
	"OutTiktok/apps/publish/publish"
	"context"
	"github.com/jinzhu/copier"
	"math"
	"strconv"

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
	stop := in.LatestTime - 1
	if in.LatestTime == 0 {
		stop = math.MaxInt64
	}

	// 查询缓存
	pairs, err := l.svcCtx.Redis.ZrevrangebyscoreWithScoresAndLimit("feed", 0, stop, 0, 30)
	if err != nil {
		return &feed.FeedRes{Status: -1}, nil
	}
	if len(pairs) == 0 {
		return &feed.FeedRes{}, nil
	}

	videoIds := make([]int64, len(pairs))
	for i, pair := range pairs {
		id, _ := strconv.ParseInt(pair.Key, 10, 64)
		videoIds[i] = id
	}

	r, err := l.svcCtx.PublishClient.GetVideos(context.Background(), &publish.GetVideosReq{
		UserId:   in.UserId,
		VideoIds: videoIds,
	})
	if err != nil {
		return &feed.FeedRes{Status: -1}, nil
	}

	videoList := make([]*feed.Video, len(videoIds))
	for i, video := range r.VideoList {
		videoList[i] = &feed.Video{}
		_ = copier.Copy(videoList[i], &video)
	}

	return &feed.FeedRes{
		NextTime:  pairs[len(pairs)-1].Score,
		VideoList: videoList,
	}, nil
}
