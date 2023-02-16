package relation

import (
	"OutTiktok/apps/gateway/pkg/jwt"
	"OutTiktok/apps/relation/relation"
	"context"
	"github.com/jinzhu/copier"

	"OutTiktok/apps/gateway/internal/svc"
	"OutTiktok/apps/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFollowListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFollowListLogic {
	return &RelationFollowListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationFollowListLogic) RelationFollowList(req *types.RelationFollowListReq) (resp *types.RelationFollowListRes, err error) {
	resp = &types.RelationFollowListRes{}

	// 验证Token
	var ThisId int64
	claims, err := jwt.VerifyToken(req.Token)
	if err == nil {
		ThisId = claims.UserId
	}

	r, err := l.svcCtx.RelationClient.FollowList(context.Background(), &relation.FollowListReq{
		ThisId: ThisId,
		UserId: req.UserId,
	})

	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = err.Error()
		return resp, nil
	}

	if r.Status != 0 {
		resp.StatusCode = -1
		resp.StatusMsg = "获取关注列表失败"
		return
	}

	resp.UserList = make([]types.User, len(r.Users))
	for i, user := range r.Users {
		_ = copier.Copy(&resp.UserList[i], &user)
	}

	return
}
