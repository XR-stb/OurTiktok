package user

import (
	"OutTiktok/apps/gateway/pkg/jwt"
	"OutTiktok/apps/user/user"
	"context"

	"OutTiktok/apps/gateway/internal/svc"
	"OutTiktok/apps/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	resp = &types.LoginRes{}
	// 检查参数
	if req.Username == "" || req.Password == "" {
		resp.StatusCode = -1
		resp.StatusMsg = "用户名或密码为空"
		return
	}

	r, err := l.svcCtx.UserClient.Login(context.Background(), &user.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = err.Error()
		return
	}
	if r.Status != 0 {
		resp.StatusCode = -1
		resp.StatusMsg = "用户名或密码不正确"
		return
	}
	resp.UserId = r.UserId

	// 生成Token
	token := jwt.GetToken(&jwt.JWTClaims{
		UserId:   r.UserId,
		Username: req.Username,
	})
	resp.Token = token

	return
}
