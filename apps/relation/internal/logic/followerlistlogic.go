package logic

import (
	"OutTiktok/apps/relation/internal/svc"
	"OutTiktok/apps/relation/relation"
	"OutTiktok/apps/user/userclient"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
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
	userId := in.UserId
	thisId := in.ThisId

	// 查询数据库
	var followerIds []int64
	l.svcCtx.DB.Table("relations").Select("follower_id").Where("followed_id = ? AND status = ?", userId, 1).Find(&followerIds)

	// 获取用户信息
	users := make([]*relation.UserInfo, len(followerIds))
	if r, err := l.svcCtx.UserClient.GetUsers(context.Background(), &userclient.GetUsersReq{
		UserIds: followerIds,
		ThisId:  thisId,
	}); err == nil {
		for i, user := range r.Users {
			users[i] = &relation.UserInfo{
				Id:            user.Id,
				Name:          user.Username,
				FollowCount:   user.FollowCount,
				FollowerCount: user.FollowerCount,
				IsFollow:      user.IsFollow,
			}
		}
	}

	return &relation.FollowerListRes{
		Users: users,
	}, nil
}
