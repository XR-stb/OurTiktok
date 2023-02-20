package user

import (
	"OutTiktok/apps/gateway/pkg/jwt"
	"OutTiktok/apps/user/userclient"
	"context"
	"github.com/jinzhu/copier"

	"OutTiktok/apps/gateway/internal/svc"
	"OutTiktok/apps/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.UserReq) (resp *types.UserRes, err error) {
	resp = &types.UserRes{}
	// 检查参数
	if req.UserId == 0 {
		resp.StatusCode = -1
		resp.StatusMsg = "用户ID为空"
		return
	}

	// 验证Token
	var ThisId int64
	claims, err := jwt.VerifyToken(req.Token)
	if err == nil {
		ThisId = claims.UserId
	}

	// 调用RPC服务
	r, err := l.svcCtx.UserClient.User(l.ctx, &userclient.UserReq{
		UserId: req.UserId,
		ThisId: ThisId,
	})
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = err.Error()
		return
	}
	if r.Status != 0 {
		resp.StatusCode = -1
		resp.StatusMsg = "用户不存在"
		return
	}
	l.Info(r.User)

	_ = copier.Copy(&resp.User, r.User)
	return
}
