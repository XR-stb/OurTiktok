package favorite

import (
	"OutTiktok/apps/favorite/favorite"
	"OutTiktok/apps/gateway/pkg/jwt"
	"context"
	"github.com/jinzhu/copier"

	"OutTiktok/apps/gateway/internal/svc"
	"OutTiktok/apps/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteListLogic) FavoriteList(req *types.FavoriteListReq) (resp *types.FavoriteListRes, err error) {
	resp = &types.FavoriteListRes{}

	// 验证Token
	var thisId int64
	claims, err := jwt.VerifyToken(req.Token)
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = err.Error()
		return
	}
	thisId = claims.UserId

	r, err := l.svcCtx.FavoriteClient.List(context.Background(), &favorite.ListReq{
		UserId: req.UserId,
		ThisId: thisId,
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
		_ = copier.Copy(&resp.VideoList[i], r.VideoList[i])
	}

	return
}
