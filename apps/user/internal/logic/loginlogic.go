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

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginRes, error) {
	username := in.Username
	password := in.Password

	// 查询账号并验证密码
	u := dao.User{}
	password = fmt.Sprintf("%x", md5.Sum([]byte(password)))
	if err := l.svcCtx.DB.Where("username=? AND password=?", username, password).First(&u).Error; err != nil {
		return &user.LoginRes{
			Status: -1,
		}, nil
	}

	// 写入缓存
	key := fmt.Sprintf("uinfo_%d", u.Id)
	val := fmt.Sprintf("%s_%s_%s_%s", u.Username, u.Avatar, u.BackgroundImage, u.Signature)
	_ = l.svcCtx.Redis.Setex(key, val, 86400)

	return &user.LoginRes{
		UserId: u.Id,
	}, nil
}
