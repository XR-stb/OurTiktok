package logic

import (
	"OutTiktok/apps/relation/internal/svc"
	"OutTiktok/apps/relation/relation"
	"OutTiktok/dao"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetRelationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRelationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRelationsLogic {
	return &GetRelationsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRelationsLogic) GetRelations(in *relation.GetRelationsReq) (*relation.GetRelationsRes, error) {
	thisId := in.ThisId
	userIds := in.UserIds

	// 查询数据库
	users := make([]*relation.UserInfo, len(userIds))
	for i, uid := range userIds {
		relationModel := dao.Relation{}
		if thisId != uid {
			l.svcCtx.DB.Where("followed_id = ? AND follower_id = ? AND status = ?", uid, thisId, 1).Find(&relationModel)
		}
		var followCount int64
		l.svcCtx.DB.Model(&relationModel).Where("follower_id = ? AND status = ?", uid, 1).Count(&followCount)
		var followerCount int64
		l.svcCtx.DB.Model(&relationModel).Where("followed_id = ? AND status = ?", uid, 1).Count(&followerCount)
		users[i] = &relation.UserInfo{
			IsFollow:      isFollow(relationModel.Status),
			FollowCount:   followCount,
			FollowerCount: followerCount,
		}
	}

	return &relation.GetRelationsRes{
		Users: users,
	}, nil
}

func isFollow(i int32) bool {
	if i == 1 {
		return true
	}
	return false
}
