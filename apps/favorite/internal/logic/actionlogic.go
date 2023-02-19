package logic

import (
	"OutTiktok/dao"
	"context"
	"fmt"
	"strconv"

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
	// 更新数据库
	favoriteModel := dao.Favorite{
		VideoId: in.VideoId,
		UserId:  in.UserId,
		Status:  in.ActionType,
	}

	var err error
	if in.ActionType == 1 {
		if err := l.svcCtx.DB.Create(&favoriteModel).Error; err != nil {
			err = l.svcCtx.DB.Model(&favoriteModel).Where("video_id=? AND user_id=?", in.VideoId, in.UserId).Update("status", in.ActionType).Error
		}
	} else {
		err = l.svcCtx.DB.Model(&favoriteModel).Where("video_id=? AND user_id=?", in.VideoId, in.UserId).Update("status", in.ActionType).Error
	}
	if err != nil {
		return &favorite.ActionRes{Status: -1}, err
	}

	// 更新缓存
	key := fmt.Sprintf("fc_%d", in.VideoId)
	key2 := fmt.Sprintf("fv_%d", in.UserId)
	res, err := l.svcCtx.Redis.Get(key)
	if err != nil {
		return &favorite.ActionRes{}, err
	}

	if res == "" {
		var count int64
		l.svcCtx.DB.Model(&favoriteModel).Where("video_id=?", in.VideoId).Count(&count)
		_ = l.svcCtx.Redis.Setex(key, strconv.FormatInt(count, 10), 86400)
	} else {
		if in.ActionType == 1 {
			_, _ = l.svcCtx.Redis.Incr(key)
		} else {
			_, _ = l.svcCtx.Redis.Decr(key)
		}
		_ = l.svcCtx.Redis.Expire(key, 86400)
	}

	if in.ActionType == 1 {
		_, _ = l.svcCtx.Redis.Sadd(key2, 0, in.VideoId) // 0占位
	} else {
		_, _ = l.svcCtx.Redis.Srem(key2, in.VideoId)
	}
	_ = l.svcCtx.Redis.Expire(key2, 86400)

	return &favorite.ActionRes{}, nil
}
