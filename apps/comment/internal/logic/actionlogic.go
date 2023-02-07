package logic

import (
	"context"

	"OutTiktok/apps/comment/comment"
	"OutTiktok/apps/comment/internal/svc"

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

func (l *ActionLogic) Action(in *comment.ActionReq) (*comment.ActionRes, error) {
	// todo: add your logic here and delete this line

	return &comment.ActionRes{}, nil
}
