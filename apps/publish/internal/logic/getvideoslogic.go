package logic

import (
	"OutTiktok/apps/comment/comment"
	"OutTiktok/apps/favorite/favoriteclient"
	"OutTiktok/apps/user/userclient"
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"strconv"
	"strings"

	"OutTiktok/apps/publish/internal/svc"
	"OutTiktok/apps/publish/publish"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideosLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideosLogic {
	return &GetVideosLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideosLogic) GetVideos(in *publish.GetVideosReq) (*publish.GetVideosRes, error) {
	videoIds := in.VideoIds

	// 查询缓存
	nonCacheList := make([]int64, 0, len(videoIds)) // 未命中列表
	for _, id := range in.VideoIds {
		// 从缓存中查询
		key := fmt.Sprintf("vinfo_%d", id)
		result, err := l.svcCtx.Redis.Get(key)
		if err != nil || result == "" {
			// 加入未命中列表
			nonCacheList = append(nonCacheList, id)
			continue
		}
		// 刷新过期时间
		_ = l.svcCtx.Redis.Expire(key, 86400)

		//l.svcCtx.VideoCache[id] = parseToVideo(id, result)
		l.svcCtx.VideoCache.Store(id, parseToVideo(id, result))
	}

	// 查询数据库
	if len(nonCacheList) > 0 {
		var queryVideoList []*publish.Video
		l.svcCtx.DB.Where("id IN ?", nonCacheList).Find(&queryVideoList)
		for _, video := range queryVideoList {
			//l.svcCtx.VideoCache[video.Id] = video
			l.svcCtx.VideoCache.Store(video.Id, video)
			// 写回数据库
			key := fmt.Sprintf("vinfo_%d", video.Id)
			val := fmt.Sprintf("%d_%s_%s_%s", video.AuthorId, video.PlayUrl, video.CoverUrl, video.Title)
			_ = l.svcCtx.Redis.Setex(key, val, 86400)
		}
	}

	// 写入结果,并获取作者ID
	videoList := make([]*publish.Video, len(videoIds))
	AuthorIds := make([]int64, 0, len(videoIds))
	for i, id := range videoIds {
		//v := l.svcCtx.VideoCache[id]
		val, _ := l.svcCtx.VideoCache.Load(id)
		v := val.(*publish.Video) //转型一下
		videoList[i] = v
		AuthorIds = append(AuthorIds, v.AuthorId)
	}

	// 查询作者信息
	for _, video := range videoList {
		AuthorIds = append(AuthorIds, video.AuthorId)
	}
	if r, err := l.svcCtx.UserClient.GetUsers(context.Background(), &userclient.GetUsersReq{
		ThisId:  in.UserId,
		UserIds: AuthorIds,
	}); err == nil {
		for _, user := range r.Users {
			//l.svcCtx.UserCache[user.Id] = user
			l.svcCtx.UserCache.Store(user.Id, user)
		}
		for i := 0; i < len(videoList); i++ {
			//userinfo := l.svcCtx.UserCache[videoList[i].AuthorId]
			userinfo, _ := l.svcCtx.UserCache.Load(videoList[i].AuthorId)
			videoList[i].Author = &publish.UserInfo{}
			_ = copier.Copy(videoList[i].Author, userinfo)
		}
	}

	// 查询点赞信息
	if r, err := l.svcCtx.FavoriteClient.GetVideoFavorite(context.Background(), &favoriteclient.GetVideoFavoriteReq{
		UserId:      in.UserId,
		VideoIds:    videoIds,
		AllFavorite: in.AllFavorite,
	}); err == nil {
		for i := 0; i < len(videoList); i++ {
			videoList[i].FavoriteCount = r.Favorites[i].FavoriteCount
			videoList[i].IsFavorite = r.Favorites[i].IsFavorite
		}
	}

	// 查询评论数量
	if r, err := l.svcCtx.CommentClient.GetCommentCount(context.Background(), &comment.GetCommentCountReq{VideoIds: videoIds}); err == nil {
		for i, count := range r.Counts {
			videoList[i].CommentCount = count
		}
	}

	return &publish.GetVideosRes{
		VideoList: videoList,
	}, nil
}

// 解析成结构体
func parseToVideo(id int64, str string) *publish.Video {
	splits := strings.Split(str, "_")
	authorId, _ := strconv.ParseInt(splits[0], 10, 64)
	return &publish.Video{
		Id:       id,
		AuthorId: authorId,
		PlayUrl:  splits[1],
		CoverUrl: splits[2],
		Title:    splits[3],
	}
}
