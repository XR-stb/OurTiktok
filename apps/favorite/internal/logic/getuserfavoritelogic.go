package logic

import (
	"OutTiktok/apps/publish/publish"
	"context"
	"fmt"
	"strconv"

	"OutTiktok/apps/favorite/favorite"
	"OutTiktok/apps/favorite/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFavoriteLogic {
	return &GetUserFavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserFavoriteLogic) GetUserFavorite(in *favorite.GetUserFavoriteReq) (*favorite.GetUserFavoriteRes, error) {
	resList := make([]*favorite.UserFavorite, len(in.Users))

	for i, id := range in.Users {
		// 查询缓存
		key := fmt.Sprintf("fv_%d", id)
		count, err := l.svcCtx.Redis.Scard(key)
		if err != nil || count == 0 {
			// 查询数据库
			var videoIds []interface{}
			rows := l.svcCtx.DB.Table("favorites").Select("video_id").Where("user_id = ? AND status = ?", id, 1).Find(&videoIds).RowsAffected
			_, _ = l.svcCtx.Redis.Sadd(key, append(videoIds, 0))
			resList[i] = &favorite.UserFavorite{FavoriteCount: rows}
		} else {
			resList[i] = &favorite.UserFavorite{FavoriteCount: count - 1}
		}
		_ = l.svcCtx.Redis.Expire(key, 86400)

		// 获取用户发布的视频ID
		r, err := l.svcCtx.PublishClient.GetVideoIds(context.Background(), &publish.GetVideoIdsReq{UserId: id})
		if err != nil {
			continue
		}
		// 查询视频获赞
		var totalFavorited int64
		for _, vid := range r.VideoIds {
			// 查询缓存
			key := fmt.Sprintf("fc_%d", vid)
			count, _ := l.svcCtx.Redis.Get(key)
			if count == "" {
				// 查询数据库
				var sqlCount int64
				l.svcCtx.DB.Table("favorites").Where("video_id = ? AND status = ?", vid, 1).Count(&sqlCount)
				_ = l.svcCtx.Redis.Setex(key, strconv.FormatInt(sqlCount, 10), 86400)
				totalFavorited += sqlCount
			} else {
				count, _ := strconv.ParseInt(count, 10, 64)
				totalFavorited += count
			}
			_ = l.svcCtx.Redis.Expire(key, 86400)
		}
		resList[i].TotalFavorited = totalFavorited
	}

	return &favorite.GetUserFavoriteRes{
		Favorites: resList,
	}, nil
}
