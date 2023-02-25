package logic

import (
	"context"
	"fmt"

	"OutTiktok/apps/publish/internal/svc"
	"OutTiktok/apps/publish/publish"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWorkCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetWorkCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWorkCountLogic {
	return &GetWorkCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetWorkCountLogic) GetWorkCount(in *publish.GetWorkCountReq) (*publish.GetWorkCountRes, error) {
	counts := make([]int64, len(in.UserId))

	for i, id := range in.UserId {
		// 从缓存中查询
		key := fmt.Sprintf("uv_%d", id)
		count, err := l.svcCtx.Redis.Scard(key)
		if err != nil || count == 0 { // 未命中
			// 从数据库中查询
			var videoIds []interface{}
			l.svcCtx.DB.Table("videos").Select("id").Where("author_id = ?", id).Find(&videoIds)

			// 写回缓存
			_, _ = l.svcCtx.Redis.Sadd(key, append(videoIds, 0))
			_ = l.svcCtx.Redis.Expire(key, 86400)
			counts[i] = int64(len(videoIds))
		} else {
			_ = l.svcCtx.Redis.Expire(key, 86400)
			counts[i] = count - 1
		}
	}

	return &publish.GetWorkCountRes{
		Counts: counts,
	}, nil
}
