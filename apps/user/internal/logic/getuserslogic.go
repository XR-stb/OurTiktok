package logic

import (
	"context"

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
	userIds := in.UserIds

	// 查询数据库
	var users []*user.UserInfo
	l.svcCtx.DB.Table("users").Where("id IN ?", userIds).Find(&users)

	return &user.GetUsersRes{
		Users: users,
	}, nil
}
