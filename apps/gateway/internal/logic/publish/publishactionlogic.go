package publish

import (
	"OutTiktok/apps/gateway/pkg/jwt"
	"OutTiktok/apps/publish/publish"
	"context"

	"OutTiktok/apps/gateway/internal/svc"
	"OutTiktok/apps/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishActionLogic {
	return &PublishActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishActionLogic) PublishAction(req *types.PublishActionReq) (resp *types.PublishActionRes, err error) {
	resp = &types.PublishActionRes{}
	// 验证Token
	var UserId int64
	claims, err := jwt.VerifyToken(req.Token)
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = err.Error()
		return
	}
	UserId = claims.UserId

	if req.Data == nil || len(req.Data) == 0 {
		resp.StatusCode = -1
		resp.StatusMsg = "Empty data"
		return
	}

	r, err := l.svcCtx.PublishClient.Action(context.Background(), &publish.ActionReq{
		Data:   req.Data,
		UserId: UserId,
		Title:  req.Title,
	})
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = err.Error()
		return
	}
	if r.Status != 0 {
		resp.StatusCode = -1
		resp.StatusMsg = "上传失败"
		return
	}

	return
}
