package logic

import (
	"OutTiktok/apps/favorite/favorite"
	"OutTiktok/apps/favorite/internal/svc"
	"OutTiktok/apps/publish/publishclient"
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"strconv"

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

func (l *ListLogic) List(in *favorite.ListReq) (*favorite.ListRes, error) {
	var videoIds []int64

	// 查询缓存
	key := fmt.Sprintf("fv_%d", in.UserId)
	result, err := l.svcCtx.Redis.Smembers(key)
	if err != nil || len(result) == 0 { // 未命中
		// 查询数据库
		l.svcCtx.DB.Table("favorites").Select("video_id").Where("user_id = ?", in.UserId).Find(&videoIds)
		// 写回数据库
		temp := make([]interface{}, len(videoIds), len(videoIds)+1)
		for i, id := range videoIds {
			temp[i] = id
		}
		_, _ = l.svcCtx.Redis.Sadd(key, append(temp, 0))
		_ = l.svcCtx.Redis.Expire(key, 86400)
	} else if len(result) == 1 { // 命中为空
		_ = l.svcCtx.Redis.Expire(key, 86400)
		return &favorite.ListRes{}, nil
	} else { // 命中
		_ = l.svcCtx.Redis.Expire(key, 86400)
		videoIds = make([]int64, 0, len(result)-1)
		for _, id := range result {
			if id == "0" {
				continue
			}
			id, _ := strconv.ParseInt(id, 10, 64)
			videoIds = append(videoIds, id)
		}
	}

	if len(videoIds) == 0 {
		return &favorite.ListRes{}, nil
	}

	// 查询视频信息
	var videoList []*favorite.Video
	if r, err := l.svcCtx.PublishClient.GetVideos(context.Background(), &publishclient.GetVideosReq{
		UserId:      in.ThisId,
		VideoIds:    videoIds,
		AllFavorite: true,
	}); err == nil {
		videoList = make([]*favorite.Video, len(videoIds))
		for i, v := range r.VideoList {
			videoList[i] = &favorite.Video{}
			_ = copier.Copy(videoList[i], &v)
		}
	}

	return &favorite.ListRes{
		VideoList: videoList,
	}, nil
}
