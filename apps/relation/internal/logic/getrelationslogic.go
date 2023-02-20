package logic

import (
	"OutTiktok/apps/relation/internal/svc"
	"OutTiktok/apps/relation/relation"
	"context"
	"fmt"
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
	key0 := fmt.Sprintf("follow_%d", in.ThisId)
	if ttl, _ := l.svcCtx.Redis.Ttl(key0); ttl < 0 {
		var followIds []int64
		l.svcCtx.DB.Table("relations").Select("followed_id").Where("follower_id = ? AND status = ?", in.ThisId, 1).Find(&followIds)
		_, _ = l.svcCtx.Redis.Sadd(key0, append(followIds, 0))
	}
	_ = l.svcCtx.Redis.Expire(key0, 86400)
	resList := make([]*relation.UserRelation, len(in.UserIds))
	// 查询缓存
	for i, id := range in.UserIds {
		resList[i] = &relation.UserRelation{}
		// 关注数量
		key := fmt.Sprintf("follow_%d", id)
		count, err := l.svcCtx.Redis.Scard(key)
		if err != nil || count == 0 {
			// 查询数据库
			var followIds []int64
			l.svcCtx.DB.Table("relations").Select("followed_id").Where("follower_id = ? AND status = ?", id, 1).Find(&followIds)
			_, _ = l.svcCtx.Redis.Sadd(key, append(followIds, 0))
			_ = l.svcCtx.Redis.Expire(key, 86400)
			resList[i].FollowCount = int64(len(followIds))
		} else {
			resList[i].FollowCount = count
		}

		// 粉丝数量
		key = fmt.Sprintf("fans_%d", id)
		count, err = l.svcCtx.Redis.Scard(key)
		if err != nil || count == 0 {
			// 查询数据库
			var followerIds []int64
			l.svcCtx.DB.Table("relations").Select("follower_id").Where("followed_id = ? AND status = ?", id, 1).Find(&followerIds)
			_, _ = l.svcCtx.Redis.Sadd(key, append(followerIds, 0))
			_ = l.svcCtx.Redis.Expire(key, 86400)
			resList[i].FollowerCount = int64(len(followerIds))
		} else {
			resList[i].FollowerCount = count
		}

		if in.AllFollow {
			resList[i].IsFollow = true
		} else {
			is, _ := l.svcCtx.Redis.Sismember(key0, id)
			resList[i].IsFollow = is
		}
	}
	return &relation.GetRelationsRes{
		Relations: resList,
	}, nil
}
