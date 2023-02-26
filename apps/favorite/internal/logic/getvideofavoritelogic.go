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

	// 检查用户点赞列表是否缓存
	key := fmt.Sprintf("fv_%d", in.UserId)
	if caching, _ := l.svcCtx.Redis.Sismember(key, 0); !caching {
		var videoIds []interface{}
		l.svcCtx.DB.Table("favorites").Select("video_id").Where("user_id = ? AND status = ?", in.UserId, 1).Find(&videoIds)
		_, _ = l.svcCtx.Redis.Sadd(key, append(videoIds, 0))
	}
	_ = l.svcCtx.Redis.Expire(key, 86400)

	// 查询缓存
	for i, vid := range in.VideoIds {
		// 查询视频点赞数量
		var count int64
		key2 := fmt.Sprintf("fc_%d", vid)
		val, _ := l.svcCtx.Redis.Get(key2)
		if val == "" { // 未命中
			// 查询数据库
			l.svcCtx.DB.Table("favorites").Where("video_id = ? AND status = ?", vid, 1).Count(&count)
			// 写回缓存
			_ = l.svcCtx.Redis.Setex(key2, strconv.FormatInt(count, 10), 86400)
		}
		// 刷新过期时间
		_ = l.svcCtx.Redis.Expire(key2, 86400)

		count, _ = strconv.ParseInt(val, 10, 64)
		resList[i] = &favorite.VideoFavorite{FavoriteCount: count}

		// 查询用户是否点赞视频
		if in.AllFavorite {
			resList[i].IsFavorite = true
		} else {
			is, _ := l.svcCtx.Redis.Sismember(key, vid)
			resList[i].IsFavorite = is
		}
	}

	return &favorite.GetVideoFavoriteRes{
		Favorites: resList,
	}, nil
}
