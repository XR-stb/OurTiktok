package logic

import (
	"OutTiktok/apps/favorite/favorite"
	"OutTiktok/apps/publish/publish"
	"context"
	"fmt"
	"strings"

	"OutTiktok/apps/user/internal/svc"
	"OutTiktok/apps/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserLogic) User(in *user.UserReq) (*user.UserRes, error) {
	var u *user.UserInfo
	// 查询缓存
	key := fmt.Sprintf("uinfo_%d", in.UserId)
	val, err := l.svcCtx.Redis.Get(key)
	if err == nil && val != "" {
		u = parseToUser(in.UserId, val)
	} else {
		// 查询数据库
		u = &user.UserInfo{}
		result := l.svcCtx.DB.Table("users").Where("id=?", in.UserId).First(u)
		if result.Error != nil || result.RowsAffected < 1 {
			return &user.UserRes{Status: -1}, nil
		}
	}

	// 获取点赞信息
	if r, err := l.svcCtx.FavoriteClient.GetUserFavorite(context.Background(), &favorite.GetUserFavoriteReq{Users: []int64{in.UserId}}); err == nil {
		u.FavoriteCount = r.Favorites[0].FavoriteCount
		u.TotalFavorited = r.Favorites[0].TotalFavorited
	}

	// 获取发布数量
	if r, err := l.svcCtx.PublishClient.GetWorkCount(context.Background(), &publish.GetWorkCountReq{UserId: []int64{in.UserId}}); err == nil {
		u.WorkCount = r.Counts[0]
	}

	// TODO: 获取关注信息

	return &user.UserRes{
		User: u,
	}, nil
}

func parseToUser(id int64, str string) *user.UserInfo {
	splits := strings.Split(str, "_")
	return &user.UserInfo{
		Id:              id,
		Username:        splits[0],
		Avatar:          splits[1],
		BackgroundImage: splits[2],
		Signature:       splits[3],
	}
}
