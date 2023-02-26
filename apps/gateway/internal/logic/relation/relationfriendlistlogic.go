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

type RelationFriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFriendListLogic {
	return &RelationFriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationFriendListLogic) RelationFriendList(req *types.RelationFriendListReq) (resp *types.RelationFriendListRes, err error) {
	resp = &types.RelationFriendListRes{}

	// 验证Token
	var ThisId int64
	claims, err := jwt.VerifyToken(req.Token)
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = err.Error()
		return
	}
	ThisId = claims.UserId

	r, err := l.svcCtx.RelationClient.FriendList(context.Background(), &relation.FriendListReq{
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
		resp.StatusMsg = "获取朋友列表失败"
		return
	}

	resp.UserList = make([]types.FriendUser, len(r.Users))
	for i, user := range r.Users {
		_ = copier.Copy(&resp.UserList[i], &user)
	}

	return
}
