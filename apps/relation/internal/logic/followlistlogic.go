package logic

import (
	"OutTiktok/apps/relation/internal/svc"
	"OutTiktok/apps/relation/relation"
	"OutTiktok/apps/user/userclient"
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

type FollowListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowListLogic {
	return &FollowListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FollowList 获取关注列表
func (l *FollowListLogic) FollowList(in *relation.FollowListReq) (*relation.FollowListRes, error) {
	userId := in.UserId
	thisId := in.ThisId
	var followIds []int64

	// 查询缓存
	key := fmt.Sprintf("follow_%d", userId)
	result, err := l.svcCtx.Redis.Smembers(key)
	if err != nil || len(result) == 0 {
		// 查询数据库
		l.svcCtx.DB.Table("relations").Select("followed_id").Where("follower_id = ? AND status = ?", userId, 1).Find(&followIds)
		_, _ = l.svcCtx.Redis.Sadd(key, append(followIds, 0))
		_ = l.svcCtx.Redis.Expire(key, 86400)
	} else {
		_ = l.svcCtx.Redis.Expire(key, 86400)
		followIds = make([]int64, 0, len(result))
		for _, id := range result {
			if id == "0" {
				continue
			}
			id, _ := strconv.ParseInt(id, 10, 64)
			followIds = append(followIds, id)
		}
	}
	if len(followIds) < 1 {
		return &relation.FollowListRes{}, nil
	}

	// 获取用户信息
	users := make([]*relation.UserInfo, len(followIds))
	if r, err := l.svcCtx.UserClient.GetUsers(context.Background(), &userclient.GetUsersReq{
		UserIds: followIds,
		ThisId:  thisId,
	}); err == nil {
		for i, user := range r.Users {
			users[i] = &relation.UserInfo{}
			_ = copier.Copy(users[i], &user)
		}
	}

	return &relation.FollowListRes{
		Users: users,
	}, nil
}
