package logic

import (
	"OutTiktok/apps/comment/comment"
	"OutTiktok/apps/favorite/favoriteclient"
	"OutTiktok/apps/user/userclient"
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"strconv"

	"OutTiktok/apps/publish/internal/svc"
	"OutTiktok/apps/publish/publish"

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

func (l *ListLogic) List(in *publish.ListReq) (*publish.ListRes, error) {
	authorId := in.UserId
	var videoList []*publish.Video

	// 查询缓存
	key := fmt.Sprintf("uv_%d", in.UserId)
	result, err := l.svcCtx.Redis.Smembers(key)
	videoIds := make([]int64, 0, len(result))
	if err != nil || len(result) == 0 { // 未命中
		// 查询数据库
		l.svcCtx.DB.Where("author_id=?", authorId).Order("id desc").Find(&videoList)
		// 写回缓存
		videoIds := make([]interface{}, len(videoList))
		videoIds[0] = 0
		for i, video := range videoList {
			videoIds[i+1] = video.Id
		}
		_, _ = l.svcCtx.Redis.Sadd(key, videoIds...)
		_ = l.svcCtx.Redis.Expire(key, 86400)
		if len(videoList) == 0 {
			return &publish.ListRes{}, nil
		}
	} else if len(result) == 1 { // 命中但为空
		_ = l.svcCtx.Redis.Expire(key, 86400)
		return &publish.ListRes{}, nil
	} else { // 命中
		_ = l.svcCtx.Redis.Expire(key, 86400)
		for i := len(result) - 1; i >= 0; i-- {
			if result[i] == "0" {
				continue
			}
			id, _ := strconv.ParseInt(result[i], 10, 64)
			videoIds = append(videoIds, id)
		}

		nonCacheList := make([]int64, 0, len(result)) // 未命中列表
		// 查询视频信息
		for _, id := range videoIds {
			key := fmt.Sprintf("vinfo_%d", id)
			str, err := l.svcCtx.Redis.Get(key)
			if err != nil || str == "" { // 未命中
				nonCacheList = append(nonCacheList, id)
				continue
			}
			_ = l.svcCtx.Redis.Expire(key, 86400)

			l.svcCtx.VideoCache[id] = parseToVideo(id, str)
		}

		// 查询数据库
		if len(nonCacheList) > 0 {
			var queryVideoList []*publish.Video
			l.svcCtx.DB.Where("id IN ?", nonCacheList).Find(&queryVideoList)
			for _, video := range queryVideoList {
				l.svcCtx.VideoCache[video.Id] = video
				// 写回数据库
				key := fmt.Sprintf("vinfo_%d", video.Id)
				val := fmt.Sprintf("%d_%s_%s_%s", video.AuthorId, video.PlayUrl, video.CoverUrl, video.Title)
				_ = l.svcCtx.Redis.Setex(key, val, 86400)
			}
		}

		videoList = make([]*publish.Video, len(videoIds))
		for i, id := range videoIds {
			videoList[i] = l.svcCtx.VideoCache[id]
		}
	}

	// 查询作者信息
	if r, err := l.svcCtx.UserClient.GetUsers(context.Background(), &userclient.GetUsersReq{
		ThisId:  in.ThisId,
		UserIds: []int64{authorId},
	}); err == nil {
		userinfo := &publish.UserInfo{
			Id: authorId,
		}
		_ = copier.Copy(userinfo, r.Users[0])
		for i := 0; i < len(videoList); i++ {
			videoList[i].Author = userinfo
		}
	}

	// 查询点赞信息
	if r, err := l.svcCtx.FavoriteClient.GetFavorites(context.Background(), &favoriteclient.GetFavoritesReq{
		UserId:   in.ThisId,
		VideoIds: videoIds,
	}); err == nil {
		for i := 0; i < len(videoList); i++ {
			videoList[i].FavoriteCount = r.Favorites[i].FavoriteCount
			videoList[i].IsFavorite = r.Favorites[i].IsFavorite
		}
	}

	// 查询评论数量
	if r, err := l.svcCtx.CommentClient.GetCommentCount(context.Background(), &comment.GetCommentCountReq{VideoIds: videoIds}); err != nil {
		for i, count := range r.Counts {
			videoList[i].CommentCount = count
		}
	}

	return &publish.ListRes{
		VideoList: videoList,
	}, nil
}
