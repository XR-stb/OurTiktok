package logic

import (
	"OutTiktok/dao"
	"context"
	"fmt"

	"OutTiktok/apps/favorite/favorite"
	"OutTiktok/apps/favorite/internal/svc"

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

func (l *ActionLogic) Action(in *favorite.ActionReq) (*favorite.ActionRes, error) {
	userId := in.UserId
	videoId := in.VideoId
	actionType := in.ActionType

	// 更新数据库
	favoriteModel := dao.Favorite{
		VideoId: videoId, UserId: userId, Status: actionType,
	}

	var err error
	if actionType == 1 {
		if err := l.svcCtx.DB.Create(&favoriteModel).Error; err != nil {
			err = l.svcCtx.DB.Model(&favoriteModel).Where("video_id=? AND user_id=?", videoId, userId).Update("status", actionType).Error
		}
	} else {
		err = l.svcCtx.DB.Model(&favoriteModel).Where("video_id=? AND user_id=?", videoId, userId).Update("status", actionType).Error
	}
	if err != nil {
		l.Error(err)
		return &favorite.ActionRes{
			Status: -1,
		}, err
	}

	// 更新缓存
	key := fmt.Sprintf("fc_%d", videoId)
	res, _ := l.svcCtx.Redis.Get(key)

	if res == "" {
		var count int64
		l.svcCtx.DB.Model(&favoriteModel).Where("video_id=?", videoId).Count(&count)
		_ = l.svcCtx.Redis.Setex(key, fmt.Sprintf("%d", count), 86400)
	} else {
		if actionType == 1 {
			_, _ = l.svcCtx.Redis.Incr(key)
		} else {
			_, _ = l.svcCtx.Redis.Decr(key)
		}
		_ = l.svcCtx.Redis.Expire(key, 86400)
	}

	return &favorite.ActionRes{}, nil
}
