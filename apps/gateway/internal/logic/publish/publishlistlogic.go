package publish

import (
	"OutTiktok/apps/gateway/pkg/jwt"
	"OutTiktok/apps/publish/publish"
	"context"
	"github.com/jinzhu/copier"

	"OutTiktok/apps/gateway/internal/svc"
	"OutTiktok/apps/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishListLogic) PublishList(req *types.PublishListReq) (resp *types.PublishListRes, err error) {
	resp = &types.PublishListRes{}
	// 验证Token
	var ThisId int64
	claims, err := jwt.VerifyToken(req.Token)
	if err == nil {
		ThisId = claims.UserId
	}

	r, err := l.svcCtx.PublishClient.List(context.Background(), &publish.ListReq{
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
		resp.StatusMsg = "获取失败"
		return
	}

	resp.VideoList = make([]types.Video, len(r.VideoList))
	for i := 0; i < len(r.VideoList); i++ {
		resp.VideoList[i] = types.Video{}
		copier.Copy(resp.VideoList[i], r.VideoList[i])
	}

	return
}
