package logic

import (
	"context"
	"fmt"
	"strconv"

	"OutTiktok/apps/publish/internal/svc"
	"OutTiktok/apps/publish/publish"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoIdsLogic {
	return &GetVideoIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoIdsLogic) GetVideoIds(in *publish.GetVideoIdsReq) (*publish.GetVideoIdsRes, error) {
	var videoIds []int64
	key := fmt.Sprintf("uv_%d", in.UserId)
	members, err := l.svcCtx.Redis.Smembers(key)
	if err != nil || len(members) == 1 {
		return &publish.GetVideoIdsRes{}, nil
	} else if len(members) == 0 {
		l.svcCtx.DB.Table("videos").Select("video_id").Where("user_id = ?", in.UserId).Find(&videoIds)
		videoIds = append(videoIds, 0)
		_, _ = l.svcCtx.Redis.Sadd(key, videoIds)
	} else {
		videoIds = make([]int64, len(members))
		for i, id := range members {
			if id == "0" {
				continue
			}
			id, _ := strconv.ParseInt(id, 10, 64)
			videoIds[i] = id
		}
	}
	_ = l.svcCtx.Redis.Expire(key, 86400)

	return &publish.GetVideoIdsRes{
		VideoIds: videoIds,
	}, nil
}
