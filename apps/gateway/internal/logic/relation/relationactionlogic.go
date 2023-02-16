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
	var UserId int64
	claims, err := jwt.VerifyToken(req.Token)
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = err.Error()
		return
	}
	UserId = claims.UserId

	// 检查参数
	if req.ActionType != 1 && req.ActionType != 2 {
		resp.StatusCode = -1
		resp.StatusMsg = "操作类型数不对， 非关注（1）和取关（2）操作"
		return resp, nil
	}

	r, err := l.svcCtx.RelationClient.Action(context.Background(), &relation.ActionReq{
		ThisId:     UserId,
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
