package logic

import (
	"context"

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
	res := &user.UserRes{}
	userId := in.UserId

	// 查询数据库
	res.User = &user.UserInfo{}
	result := l.svcCtx.DB.Table("users").Where("id=?", userId).First(res.User)
	if result.Error != nil || result.RowsAffected < 1 {
		res.Status = -1
		return res, nil
	}

	return res, nil
}
