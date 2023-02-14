package logic

import (
	"OutTiktok/dao"
	"context"
	"crypto/md5"
	"fmt"

	"OutTiktok/apps/user/internal/svc"
	"OutTiktok/apps/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterRes, error) {
	username := in.Username
	password := in.Password

	// 将用户写入数据库
	u := dao.User{
		Username: username,
		Password: fmt.Sprintf("%x", md5.Sum([]byte(password))), // md5加密
	}
	if err := l.svcCtx.DB.Create(&u).Error; err != nil {
		return &user.RegisterRes{
			Status: -1,
		}, nil
	}

	return &user.RegisterRes{
		UserId: u.Id,
	}, nil
}
