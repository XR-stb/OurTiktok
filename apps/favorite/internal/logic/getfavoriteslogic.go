package logic

import (
	"OutTiktok/dao"
	"context"
	"fmt"
	"strconv"

	"OutTiktok/apps/favorite/favorite"
	"OutTiktok/apps/favorite/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFavoritesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFavoritesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFavoritesLogic {
	return &GetFavoritesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFavoritesLogic) GetFavorites(in *favorite.GetFavoritesReq) (*favorite.GetFavoritesRes, error) {
	userId := in.UserId
	videoIds := in.VideoIds

	favorites := make([]*favorite.FavoriteInfo, len(videoIds))
	for i, vid := range videoIds {
		// 查点赞数量
		// 查缓存
		key := fmt.Sprintf("fc_%d", vid)
		favoriteModel := dao.Favorite{}
		res, _ := l.svcCtx.Redis.Get(key)
		var count int64
		if res != "" {
			count, _ = strconv.ParseInt(res, 10, 64)
		} else {
			// 查数据库
			l.svcCtx.DB.Model(&favoriteModel).Where("video_id=?", vid).Count(&count)
			_ = l.svcCtx.Redis.Setex(key, fmt.Sprintf("%d", count), 86400)
		}

		if userId != 0 {
			l.svcCtx.DB.Where("video_id=? AND user_id=? AND status=?", vid, userId, 1).Find(&favoriteModel)
		}
		favorites[i] = &favorite.FavoriteInfo{
			FavoriteCount: count,
			IsFavorite:    isFavorite(favoriteModel.Status),
		}
	}

	return &favorite.GetFavoritesRes{
		Favorites: favorites,
	}, nil
}

func isFavorite(i int32) bool {
	if i == 1 {
		return true
	}
	return false
}
