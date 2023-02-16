package logic

import (
	"OutTiktok/apps/relation/internal/svc"
	"OutTiktok/apps/relation/relation"
	"OutTiktok/apps/user/userclient"
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
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

	// 查询数据库
	var followedIds []int64
	l.svcCtx.DB.Table("relations").Select("followed_id").Where("follower_id = ? AND status = ?", userId, 1).Find(&followedIds)

	// 获取用户信息
	users := make([]*relation.UserInfo, len(followedIds))
	if r, err := l.svcCtx.UserClient.GetUsers(context.Background(), &userclient.GetUsersReq{
		UserIds: followedIds,
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
