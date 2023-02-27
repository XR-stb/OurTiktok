package logic

import (
	"OutTiktok/apps/favorite/favorite"
	"OutTiktok/apps/favorite/internal/svc"
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

	// 更新缓存，如果不在缓存则不更新，依赖下一次查询
	// 更新视频点赞数量
	key := fmt.Sprintf("fc_%d", in.VideoId)
	if ttl, _ := l.svcCtx.Redis.Ttl(key); ttl > 0 {
		if in.ActionType == 1 {
			_, _ = l.svcCtx.Redis.Incr(key)
		} else {
			_, _ = l.svcCtx.Redis.Decr(key)
		}
		_ = l.svcCtx.Redis.Expire(key, 86400)
	}

	// 更新用户点赞视频
	key2 := fmt.Sprintf("fv_%d", in.UserId)
	if ttl, _ := l.svcCtx.Redis.Ttl(key2); ttl > 0 {
		if in.ActionType == 1 {
			_, _ = l.svcCtx.Redis.Sadd(key2, in.VideoId)
		} else {
			_, _ = l.svcCtx.Redis.Srem(key2, in.VideoId)
		}
	}

	return &favorite.ActionRes{}, nil
}
