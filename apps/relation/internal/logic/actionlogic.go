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
		if err = l.svcCtx.DB.Create(&relationModel).Error; err != nil {
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
	key := fmt.Sprintf("follow_%d", in.ThisId)
	key2 := fmt.Sprintf("fans_%d", in.UserId)
	if actionType == 1 {
		if ttl, _ := l.svcCtx.Redis.Ttl(key); ttl > 0 { // 缓存存在->添加
			_, _ = l.svcCtx.Redis.Sadd(key, in.UserId)
		}
		if ttl, _ := l.svcCtx.Redis.Ttl(key2); ttl > 0 {
			_, _ = l.svcCtx.Redis.Sadd(key2, in.ThisId)
		}
	} else {
		if ttl, _ := l.svcCtx.Redis.Ttl(key); ttl > 0 {
			_, _ = l.svcCtx.Redis.Srem(key, in.UserId)
		}
		if ttl, _ := l.svcCtx.Redis.Ttl(key2); ttl > 0 {
			_, _ = l.svcCtx.Redis.Srem(key2, in.ThisId)
		}
	}

	return &relation.ActionRes{}, nil
}
