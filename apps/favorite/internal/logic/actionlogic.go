package logic

import (
	"context"

	"OutTiktok/apps/favorite/favorite"
	"OutTiktok/apps/favorite/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActionLogic {
	return &ActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActionLogic) Action(in *favorite.ActionReq) (*favorite.ActionRes, error) {
	// todo: add your logic here and delete this line

	return &favorite.ActionRes{}, nil
}
