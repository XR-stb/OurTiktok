package logic

import (
	"OutTiktok/apps/favorite/favorite"
	"OutTiktok/apps/favorite/internal/svc"
	"context"
	"fmt"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoFavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoFavoriteLogic {
	return &GetVideoFavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoFavoriteLogic) GetVideoFavorite(in *favorite.GetVideoFavoriteReq) (*favorite.GetVideoFavoriteRes, error) {
	resList := make([]*favorite.VideoFavorite, len(in.VideoIds))
	// 查询缓存
	key := fmt.Sprintf("fv_%d", in.UserId)
	caching, _ := l.svcCtx.Redis.Sismember(key, 0)
	if !caching {
		var videoIds []int64
		l.svcCtx.DB.Table("favorites").Select("video_id").Where("user_id = ? AND status = ?", in.UserId, 1).Find(&videoIds)
		videoIds = append(videoIds, 0)
		_, _ = l.svcCtx.Redis.Sadd(key, videoIds)
	}
	nonCacheList := make([]int64, 0, len(in.VideoIds))
	for i, id := range in.VideoIds {
		key2 := fmt.Sprintf("fc_%d", id)
		val, err := l.svcCtx.Redis.Get(key2)
		if err != nil {
			nonCacheList = append(nonCacheList, id)
			continue
		}
		// 刷新过期时间
		_ = l.svcCtx.Redis.Expire(key2, 86400)

		count, _ := strconv.ParseInt(val, 10, 64)
		resList[i] = &favorite.VideoFavorite{FavoriteCount: count}

		if in.AllFavorite {
			resList[i].IsFavorite = true
		} else {
			is, _ := l.svcCtx.Redis.Sismember(key, id)
			resList[i].IsFavorite = is
		}
	}
	_ = l.svcCtx.Redis.Expire(key, 86400)

	return &favorite.GetVideoFavoriteRes{
		Favorites: resList,
	}, nil
}
