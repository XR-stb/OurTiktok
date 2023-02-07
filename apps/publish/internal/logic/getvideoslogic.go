package logic

import (
	"context"

	"OutTiktok/apps/publish/internal/svc"
	"OutTiktok/apps/publish/publish"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideosLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideosLogic {
	return &GetVideosLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideosLogic) GetVideos(in *publish.GetVideosReq) (*publish.GetVideosRes, error) {
	// todo: add your logic here and delete this line

	return &publish.GetVideosRes{}, nil
}
