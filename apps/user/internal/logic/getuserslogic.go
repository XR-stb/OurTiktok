package logic

import (
	"OutTiktok/apps/favorite/favorite"
	"OutTiktok/apps/publish/publish"
	"OutTiktok/apps/relation/relation"
	"context"
	"fmt"

	"OutTiktok/apps/user/internal/svc"
	"OutTiktok/apps/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersLogic {
	return &GetUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUsersLogic) GetUsers(in *user.GetUsersReq) (*user.GetUsersRes, error) {
	users := make([]*user.UserInfo, 0, len(in.UserIds))

	//从缓存中查询
	nonCacheList := make([]int64, 0, len(in.UserIds)) //未命中列表
	for _, id := range in.UserIds {
		key := fmt.Sprintf("uinfo_%d", id)
		val, err := l.svcCtx.Redis.Get(key)
		if err != nil {
			nonCacheList = append(nonCacheList, id)
			continue
		}
		//刷新过期时间
		_ = l.svcCtx.Redis.Expire(key, 86400)

		users = append(users, parseToUser(id, val))
	}

	// 查询数据库
	if len(nonCacheList) > 0 {
		var queryUsers []*user.UserInfo
		l.svcCtx.DB.Table("users").Where("id IN ?", nonCacheList).Find(&users)

		// 写入缓存
		for _, u := range queryUsers {
			key := fmt.Sprintf("uinfo_%d", u.Id)
			val := fmt.Sprintf("%s_%s_%s_%s", u.Username, u.Avatar, u.BackgroundImage, u.Signature)
			_ = l.svcCtx.Redis.Setex(key, val, 86400)
		}

		users = append(users, queryUsers...)
	}

	// 获取点赞信息
	if r, err := l.svcCtx.FavoriteClient.GetUserFavorite(context.Background(), &favorite.GetUserFavoriteReq{Users: in.UserIds}); err == nil {
		for i, Favorite := range r.Favorites {
			users[i].FavoriteCount = Favorite.FavoriteCount
			users[i].TotalFavorited = Favorite.TotalFavorited
		}
	}

	// 获取发布数量
	if r, err := l.svcCtx.PublishClient.GetWorkCount(context.Background(), &publish.GetWorkCountReq{UserId: in.UserIds}); err == nil {
		for i, count := range r.Counts {
			users[i].WorkCount = count
		}
	}

	// 获取关注信息
	if r, err := l.svcCtx.RelationClient.GetRelations(context.Background(), &relation.GetRelationsReq{
		ThisId:    in.ThisId,
		UserIds:   in.UserIds,
		AllFollow: in.AllFollow,
	}); err == nil {
		for i, userRelation := range r.Relations {
			users[i].FollowCount = userRelation.FollowCount
			users[i].FollowerCount = userRelation.FollowerCount
			users[i].IsFollow = userRelation.IsFollow
		}
	}

	return &user.GetUsersRes{
		Users: users,
	}, nil
}
