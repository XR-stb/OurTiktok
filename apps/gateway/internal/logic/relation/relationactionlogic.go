package relation

import (
	"OutTiktok/apps/gateway/pkg/jwt"
	"OutTiktok/apps/relation/relation"
	"context"

	"OutTiktok/apps/gateway/internal/svc"
	"OutTiktok/apps/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationActionLogic {
	return &RelationActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationActionLogic) RelationAction(req *types.RelationActionReq) (resp *types.RelationActionRes, err error) {
	resp = &types.RelationActionRes{}

	// 验证Token
	var thisId int64
	claims, err := jwt.VerifyToken(req.Token)
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = err.Error()
		return
	}
	thisId = claims.UserId

	// 检查参数
	if req.ActionType != 1 && req.ActionType != 2 {
		resp.StatusCode = -1
		resp.StatusMsg = "Wrong action_type"
		return resp, nil
	}

	if req.ToUserId == thisId {
		resp.StatusCode = -1
		resp.StatusMsg = "Can't follow yourself"
		return resp, nil
	}

	r, err := l.svcCtx.RelationClient.Action(context.Background(), &relation.ActionReq{
		ThisId:     thisId,
		UserId:     req.ToUserId,
		ActionType: req.ActionType,
	})

	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = err.Error()
		return resp, nil
	}

	if r.Status != 0 {
		resp.StatusCode = -1
		resp.StatusMsg = "关注操作失败"
		return
	}

	return
}
