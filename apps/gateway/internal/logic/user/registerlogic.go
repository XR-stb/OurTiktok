package user

import (
	"OutTiktok/apps/gateway/pkg/jwt"
	"OutTiktok/apps/user/userclient"
	"context"

	"OutTiktok/apps/gateway/internal/svc"
	"OutTiktok/apps/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterRes, err error) {
	resp = &types.RegisterRes{}
	// 检查参数
	if req.Username == "" || req.Password == "" {
		resp.StatusCode = -1
		resp.StatusMsg = "用户名或密码为空"
		return
	}

	// 调用RPC服务
	r, err := l.svcCtx.UserClient.Register(l.ctx, &userclient.RegisterReq{
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
		resp.StatusMsg = "用户名被占用"
		return
	}
	resp.UserId = r.UserId

	token := jwt.GetToken(&jwt.JWTClaims{
		UserId:   r.UserId,
		Username: req.Username,
	})
	resp.Token = token

	return
}
