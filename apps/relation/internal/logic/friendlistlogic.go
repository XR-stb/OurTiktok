package logic

import (
	"OutTiktok/apps/relation/internal/svc"
	"OutTiktok/apps/relation/relation"
	"OutTiktok/apps/user/user"
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

type FriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FriendList 获取好友列表
func (l *FriendListLogic) FriendList(in *relation.FriendListReq) (*relation.FriendListRes, error) {
	// 确认缓存
	key1 := fmt.Sprintf("follow_%d", in.UserId)
	key2 := fmt.Sprintf("fans_%d", in.UserId)

	if ttl, _ := l.svcCtx.Redis.Ttl(key1); ttl < 0 {
		// 查询数据库
		var followIds []int64
		l.svcCtx.DB.Table("relations").Select("followed_id").Where("follower_id = ? AND status = ?", in.UserId, 1).Find(&followIds)
		_, _ = l.svcCtx.Redis.Sadd(key1, append(followIds, 0))
	}
	_ = l.svcCtx.Redis.Expire(key1, 86400)

	if ttl, _ := l.svcCtx.Redis.Ttl(key2); ttl < 0 {
		// 查询数据库
		var followerIds []int64
		l.svcCtx.DB.Table("relations").Select("follower_id").Where("followed_id = ? AND status = ?", in.UserId, 1).Find(&followerIds)
		_, _ = l.svcCtx.Redis.Sadd(key2, append(followerIds, 0))
	}
	_ = l.svcCtx.Redis.Expire(key2, 86400)

	var friendIds []int64
	result, err := l.svcCtx.Redis.Sinter(key1, key2)
	if err != nil || len(result) == 0 {
		return &relation.FriendListRes{Status: -1}, nil
	} else if len(result) == 1 {
		return &relation.FriendListRes{}, nil
	} else {
		friendIds = make([]int64, 0, len(result))
		for _, id := range result {
			if id == "0" {
				continue
			}
			id, _ := strconv.ParseInt(id, 10, 64)
			friendIds = append(friendIds, id)
		}
	}

	// 获取用户信息
	if r, err := l.svcCtx.UserClient.GetUsers(context.Background(), &user.GetUsersReq{
		UserIds: friendIds,
		ThisId:  in.ThisId,
	}); err != nil {
		for _, u := range r.Users {
			l.svcCtx.UserCache[u.Id] = u
		}
	}

	users := make([]*relation.FriendUser, len(friendIds))
	for i, id := range friendIds {
		users[i] = &relation.FriendUser{}
		_ = copier.Copy(users[i], l.svcCtx.UserCache[id])
	}

	return &relation.FriendListRes{
		Users: users,
	}, nil
}
