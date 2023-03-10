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

type FollowerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowerListLogic {
	return &FollowerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FollowerList 获取粉丝列表
func (l *FollowerListLogic) FollowerList(in *relation.FollowerListReq) (*relation.FollowerListRes, error) {
	var followerIds []int64

	// 查询缓存
	key := fmt.Sprintf("fans_%d", in.UserId)
	result, err := l.svcCtx.Redis.Smembers(key)
	if err != nil || len(result) == 0 {
		// 查询数据库
		l.svcCtx.DB.Table("relations").Select("follower_id").Where("followed_id = ? AND status = ?", in.UserId, 1).Find(&followerIds)
		// 写回缓存
		temp := make([]interface{}, len(followerIds), len(followerIds)+1)
		for i, id := range followerIds {
			temp[i] = id
		}
		_, _ = l.svcCtx.Redis.Sadd(key, append(temp, 0))
		_ = l.svcCtx.Redis.Expire(key, 86400)
	} else {
		_ = l.svcCtx.Redis.Expire(key, 86400)
		followerIds = make([]int64, 0, len(result))
		for _, id := range result {
			if id == "0" {
				continue
			}
			id, _ := strconv.ParseInt(id, 10, 64)
			followerIds = append(followerIds, id)
		}
	}
	if len(followerIds) < 1 {
		return &relation.FollowerListRes{}, nil
	}

	// 获取用户信息
	users := make([]*relation.UserInfo, len(followerIds))
	if r, err := l.svcCtx.UserClient.GetUsers(context.Background(), &userclient.GetUsersReq{
		UserIds: followerIds,
		ThisId:  in.ThisId,
	}); err == nil {
		for i, user := range r.Users {
			users[i] = &relation.UserInfo{}
			_ = copier.Copy(users[i], &user)
		}
	}

	return &relation.FollowerListRes{
		Users: users,
	}, nil
}
