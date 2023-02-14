package logic

import (
	"OutTiktok/apps/relation/internal/svc"
	"OutTiktok/apps/relation/relation"
	"OutTiktok/dao"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type ActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActionLogic {
	return &ActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActionLogic) Action(in *relation.ActionReq) (*relation.ActionRes, error) {
	thisId := in.ThisId
	userId := in.UserId
	actionType := in.ActionType

	// 更新数据库
	relationModel := dao.Relation{
		FollowedId: userId,
		FollowerId: thisId,
		Status:     actionType,
	}

	var err error
	if actionType == 1 {
		if err := l.svcCtx.DB.Create(&relationModel).Error; err != nil {
			err = l.svcCtx.DB.Model(&relationModel).Where("followed_id=? AND follower_id=?", userId, thisId).Update("status", actionType).Error
		}
	} else {
		err = l.svcCtx.DB.Model(&relationModel).Where("followed_id=? AND follower_id=?", userId, thisId).Update("status", actionType).Error
	}
	if err != nil {
		l.Error(err)
		return &relation.ActionRes{
			Status: -1,
		}, err
	}

	// 更新缓存
	fanskey := fmt.Sprintf("rfansc_%d", userId)
	followkey := fmt.Sprintf("rfollowc_%d", thisId)
	fansCount, _ := l.svcCtx.Redis.Get(fanskey)
	followCount, _ := l.svcCtx.Redis.Get(followkey)
	if fansCount == "" {
		var count int64
		l.svcCtx.DB.Table("relations").Where("followed_id = ?", userId).Count(&count)
		_ = l.svcCtx.Redis.Setex(fanskey, fmt.Sprintf("%d", count), 86400)
	} else {
		if actionType == 1 {
			_, _ = l.svcCtx.Redis.Incr(fanskey)
		} else {
			_, _ = l.svcCtx.Redis.Decr(fanskey)
		}
		_ = l.svcCtx.Redis.Expire(fanskey, 86400)
	}

	if followCount == "" {
		var count int64
		l.svcCtx.DB.Table("relations").Where("follower_id = ?", thisId).Count(&count)
		_ = l.svcCtx.Redis.Setex(followkey, fmt.Sprintf("%d", count), 86400)
	} else {
		if actionType == 1 {
			_, _ = l.svcCtx.Redis.Incr(followkey)
		} else {
			_, _ = l.svcCtx.Redis.Incr(followkey)
		}
		_ = l.svcCtx.Redis.Expire(followkey, 86400)
	}

	return &relation.ActionRes{}, nil
}
