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
	if err != nil || len(members) == 0 {
		l.svcCtx.DB.Table("videos").Select("id").Where("author_id = ?", in.UserId).Find(&videoIds)
		temp := make([]interface{}, len(videoIds), len(videoIds)+1)
		for i, id := range videoIds {
			temp[i] = id
		}
		_, _ = l.svcCtx.Redis.Sadd(key, append(temp, 0))
		_ = l.svcCtx.Redis.Expire(key, 86400)
	} else if len(members) == 1 {
		_ = l.svcCtx.Redis.Expire(key, 86400)
	} else {
		_ = l.svcCtx.Redis.Expire(key, 86400)
		videoIds = make([]int64, 0, len(members)-1)
		for _, id := range members {
			if id == "0" {
				continue
			}
			id, _ := strconv.ParseInt(id, 10, 64)
			videoIds = append(videoIds, id)
		}
	}

	return &publish.GetVideoIdsRes{
		VideoIds: videoIds,
	}, nil
}
