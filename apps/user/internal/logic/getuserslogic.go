package logic

import (
	"context"
	"fmt"

	"OutTiktok/apps/user/internal/svc"
	"OutTiktok/apps/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersLogic {
	return &GetUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUsersLogic) GetUsers(in *user.GetUsersReq) (*user.GetUsersRes, error) {
	users := make([]*user.UserInfo, 0, len(in.UserIds))
	nonCacheList := make([]int64, 0, len(in.UserIds))
	for _, id := range in.UserIds {
		key := fmt.Sprintf("uinfo_%d", id)
		val, err := l.svcCtx.Redis.Get(key)
		if err != nil {
			nonCacheList = append(nonCacheList, id)
			continue
		}
		//刷新过期时间
		_ = l.svcCtx.Redis.Expire(key, 86400)

		users = append(users, parseToUser(id, val))
	}

	// 查询数据库
	if len(nonCacheList) > 0 {
		var queryUsers []*user.UserInfo
		l.svcCtx.DB.Table("users").Where("id IN ?", nonCacheList).Find(&users)

		// 写入缓存
		for _, u := range queryUsers {
			key := fmt.Sprintf("uinfo_%d", u.Id)
			val := fmt.Sprintf("%s_%s_%s_%s", u.Username, u.Avatar, u.BackgroundImage, u.Signature)
			_ = l.svcCtx.Redis.Setex(key, val, 86400)
		}

		users = append(users, queryUsers...)
	}

	//TODO: 获取点赞信息

	//TODO: 获取关注信息

	return &user.GetUsersRes{
		Users: users,
	}, nil
}
