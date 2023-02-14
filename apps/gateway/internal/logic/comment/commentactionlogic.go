package comment

import (
	"OutTiktok/apps/comment/comment"
	"OutTiktok/apps/gateway/pkg/jwt"
	"context"
	"github.com/jinzhu/copier"

	"OutTiktok/apps/gateway/internal/svc"
	"OutTiktok/apps/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentActionLogic) CommentAction(req *types.CommentActionReq) (resp *types.CommentActionRes, err error) {
	resp = &types.CommentActionRes{}

	// 验证Token
	var UserId int64
	claims, err := jwt.VerifyToken(req.Token)
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = err.Error()
		return
	}
	UserId = claims.UserId

	//// 检查参数
	if req.ActionType != 1 && req.ActionType != 2 {
		resp.StatusCode = -1
		resp.StatusMsg = "操作类型数不对， 非删除（2）和增加（1）操作"
		return resp, nil
	}

	r, err := l.svcCtx.CommentClient.Action(context.Background(), &comment.ActionReq{
		UserId:     UserId,
		VideoId:    req.VideoId,
		ActionType: req.ActionType,
		Content:    req.CommentText,
		CommentId:  req.CommentId,
	})

	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = err.Error()
		return resp, nil
	}

	if r.Status != 0 {
		resp.StatusCode = -1
		resp.StatusMsg = "评论操作失败"
		return
	}

	resp.StatusMsg = "成功"
	_ = copier.Copy(&resp.Comment, r.CommentInfo)
	return
}
